package controllers

import "ASPGo/framework"

func UserLoginController(c *framework.Context) {
	c.ISetOkStatus().IJson("ok, UserLoginController")
}
