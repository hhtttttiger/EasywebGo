package http

import (
	"ASPGo/framework"
	"ASPGo/framework/middleware"
)

// NewHttpEngine 创建了一个绑定了路由的 Web 引擎
func NewHttpEngine() (*framework.Core, error) {
	// 默认启动一个 Web 引擎
	core := framework.NewCore()
	core.Use(middleware.Recovery())

	RegisterRouter(core)

	// 返回绑定路由后的 Web 引擎
	return core, nil
}
