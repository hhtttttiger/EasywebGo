package main

import (
	asphttp "ASPGo/app/http"
	"ASPGo/framework"
	"ASPGo/framework/contract"
	"ASPGo/framework/provider/app"
	"ASPGo/framework/provider/config"
	"ASPGo/framework/provider/env"
	"ASPGo/framework/provider/kernel"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// 初始化服务容器
	container := framework.NewHadeContainer()

	// 绑定App服务提供者
	container.Bind(&app.HadeAppProvider{BaseFolder: ".\\"})
	container.Bind(&env.HadeEnvProvider{})
	container.Bind(&config.HadeConfigProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := asphttp.NewHttpEngine(); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}

	// 从服务容器中获取kernel的服务实例
	kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
	// 从kernel服务实例中获取引擎
	core := kernelService.HttpEngine()

	// 创建一个Server服务
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	// 这个goroutine是启动服务的goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
