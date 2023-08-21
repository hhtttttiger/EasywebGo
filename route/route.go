package route

import (
	"ASPGo/controllers"
	"ASPGo/framework"
)

// 注册路由规则
func RegisterRouter(core *framework.Core) {

	core.Get("/user/login", controllers.UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 动态路由
		subjectApi.Delete("/:id", controllers.SubjectDelController)
		subjectApi.Put("/:id", controllers.SubjectUpdateController)
		subjectApi.Get("/list/all", controllers.SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", controllers.SubjectNameController)
		}
	}
}
