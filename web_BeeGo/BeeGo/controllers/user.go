package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	userId := c.Ctx.Input.Param(":id")
	c.Data["User"] = "User ID: " + userId
	c.TplName = "user.tpl"
}
