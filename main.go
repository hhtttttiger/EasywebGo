package main

import (
	"ASPGo/app/provider/demo"
	"ASPGo/framework"
	"ASPGo/framework/middleware"
	"ASPGo/framework/provider/app"
	"ASPGo/route"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Recovery())

	//注册
	core.Bind(&demo.DemoServiceProvider{})
	core.Bind(&app.HadeAppProvider{BaseFolder: "/tmp"})

	route.RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	go func() {
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	// 调用Server.Shutdown graceful结束
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
