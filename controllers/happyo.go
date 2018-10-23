package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"os"
	"time"
)

type HappyoController struct {
	beego.Controller
}

type shahon struct {
	Id     int    `form:"-"`
	User   string `form:"-" json:"user"`
	Gakuin string `form:"gakuin" json:"gakuin"`
	Senmon string `form:"senmon" json:"senmon"`
	Naibun string `form:"naibun" json:"naibun"`
	Hiduke string `form:"-" json:"hiduke"`
	Kidoku string `form:"-" json:"kidoku"`
}

func (c *HappyoController) Post() {
	s := shahon{}
	if err := c.ParseForm(&s); err != nil {
		println(err.Error())
	}
	s.Hiduke = time.Now().Format("2006-01-02 15:04")
	s.User = "易班用户"
	s.Kidoku = "flase"
	jsons, err := json.Marshal(s)
	if err != nil {
		println(err.Error())
	}
	f, err := os.OpenFile("database.txt", os.O_APPEND, 0644)
	if err != nil {
		println(err.Error())
	}
	defer f.Close()
	f.Write(jsons)
	f.WriteString(",")
	c.TplName = "happyo.html"
}
