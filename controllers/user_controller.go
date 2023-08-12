package controllers

import "ASPGo/framework"

func UserLoginController(c *framework.Context) error {
	c.ISetOkStatus().IJson("ok, UserLoginController")
	return nil
}
