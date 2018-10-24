package routers

import (
	"github.com/astaxie/beego"
	"github.com/nangcr/yiBan/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/sukejuru", &controllers.SukejuruController{})
	beego.Router("/toko", &controllers.TokoController{})
	beego.Router("/happyo", &controllers.HappyoController{})
	beego.Router("/browse", &controllers.BrowseController{})
	beego.Router("/rankingu", &controllers.RankinguController{})
}
