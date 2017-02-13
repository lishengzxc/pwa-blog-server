package routers

import (
	"pwa-blog-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/list/?:page", &controllers.MainController{}, "get:GetList")
    beego.Router("/article/:id", &controllers.MainController{}, "get:GetDetail")
}
