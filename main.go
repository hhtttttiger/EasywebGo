package main

import (
	"ASPGo/app/http"
	"ASPGo/framework"
	"ASPGo/framework/provider/app"
	"ASPGo/framework/provider/kernel"
)

func main() {

	// 初始化服务容器
	container := framework.NewHadeContainer()

	// 绑定App服务提供者
	container.Bind(&app.HadeAppProvider{})
	// 后续初始化需要绑定的服务提供者...

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)

	// core := framework.NewCore()
	// core.Use(middleware.Recovery())

	// //注册
	// core.Bind(&demo.DemoServiceProvider{})
	// core.Bind(&app.HadeAppProvider{BaseFolder: "/tmp"})

	// route.RegisterRouter(core)
	// server := &http.Server{
	// 	Handler: core,
	// 	Addr:    ":8888",
	// }

	// go func() {
	// 	server.ListenAndServe()
	// }()

	// quit := make(chan os.Signal)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// <-quit

	// // 调用Server.Shutdown graceful结束
	// if err := server.Shutdown(context.Background()); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }
}
