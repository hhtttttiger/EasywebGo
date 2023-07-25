package route

import (
	"ASPGo/controllers"
	"ASPGo/framework"
)

// 注册路由规则
func RegisterRouter(core *framework.Core) {
	// 静态路由+HTTP方法匹配
	core.Get("/user/login", controllers.UserLoginController)
}
