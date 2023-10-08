package main

import (
	console "ASPGo/app/console/command"
	"ASPGo/app/http"
	"ASPGo/framework"
	"ASPGo/framework/provider/app"
	"ASPGo/framework/provider/config"
	"ASPGo/framework/provider/env"
	"ASPGo/framework/provider/kernel"
)

func main() {

	// 初始化服务容器
	container := framework.NewHadeContainer()

	// 绑定App服务提供者
	container.Bind(&app.HadeAppProvider{BaseFolder: ".\\"})
	container.Bind(&env.HadeEnvProvider{})
	container.Bind(&config.HadeConfigProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
