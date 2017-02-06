package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Lishengzxc"] = "lishengzxc"
	c.TplName = "index.tpl"
}

func (c *MainController) GetList() {
	c.Ctx.WriteString("get list...")
}

func (c *MainController) GetDetail() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString(id)
}
