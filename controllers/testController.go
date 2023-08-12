package controllers

import "ASPGo/framework"

func FooControllerHandler(ctx *framework.Context) error {
	ctx.ISetOkStatus().IJson(map[string]interface{}{
		"code": 0,
	})
	return nil
}
