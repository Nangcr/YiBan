package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type RankinguController struct {
	beego.Controller
}

func (c *RankinguController) Get() {
	f, err := os.Open("database2.txt")
	if err != nil {
		println(err.Error())
	}
	defer f.Close()
	bytes, _ := ioutil.ReadAll(f)
	jsons := strings.Replace(string(bytes), "\n", "", 1)
	c.Data["Ranks"] = template.JS(jsons)
	c.TplName = "rankingu.html"
}
