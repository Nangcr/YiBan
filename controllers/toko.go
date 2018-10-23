package controllers

import (
	"github.com/astaxie/beego"
)

type TokoController struct {
	beego.Controller
}

func (c *TokoController) Get() {
	c.TplName = "toko.html"
}
