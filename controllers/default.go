package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "401.moe"
	c.Data["Email"] = "nangcr@gmail.com"
	c.TplName = "index.html"
}
