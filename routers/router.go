package routers

import (
	"pwa-blog-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/getList", &controllers.MainController{}, "get:GetList")
    beego.Router("/getDetail/:id", &controllers.MainController{}, "get:GetDetail")
}
