package controllers

import "github.com/astaxie/beego"

type SukejuruController struct {
	beego.Controller
}

func (c *SukejuruController) Get() {
	c.TplName = "sukejuru.html"
}
