package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type AboutController struct {
	web.Controller
}

func (c *AboutController) Get() {
	c.Data["Message"] = "First BeeGo app by cooler-SAI"
	c.TplName = "about.tpl"
}
