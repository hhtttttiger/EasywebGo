package controllers

import "ASPGo/framework"

func FooControllerHandler(ctx *framework.Context) error {
	ctx.SetOkStatus().Json(map[string]interface{}{
		"code": 0,
	})
	return nil
}
