package controllers

import (
	"ASPGo/framework"
	"ASPGo/framework/contract"
	"fmt"
)

func SubjectAddController(c *framework.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")
}

// 对应路由 /subject/list/all
func SubjectListController(c *framework.Context) {
	// 获取demo服务实例
	//demoService := c.MustMake(demo.Key).(demo.Service)

	// 调用服务实例的方法
	//foo := demoService.GetFoo()

	// 输出结果
	//c.ISetOkStatus().IJson(foo)

	//获取password
	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	password := configService.GetString("database.mysql.password")
	// 打印出来
	c.ISetOkStatus().IJson(password)
}

func SubjectDelController(c *framework.Context) {
	c.ISetOkStatus().IJson("ok, SubjectDelController")
}

func SubjectUpdateController(c *framework.Context) {
	c.ISetOkStatus().IJson("ok, SubjectUpdateController")
}

func SubjectGetController(c *framework.Context) {
	subjectId, _ := c.DefaultParamInt("id", 0)
	c.ISetOkStatus().IJson("ok, SubjectGetController:" + fmt.Sprint(subjectId))

}

func SubjectNameController(c *framework.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")
}
