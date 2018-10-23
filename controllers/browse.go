package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type BrowseController struct {
	beego.Controller
}

func (c *BrowseController) Get() {
	f, err := os.Open("database.txt")
	if err != nil {
		println(err.Error())
	}
	defer f.Close()
	bytes, _ := ioutil.ReadAll(f)
	jsons := strings.Replace(string(bytes), "\n", "", 1)
	c.Data["Shahons"] = template.JS(jsons)
	c.TplName = "browse.html"
}
