package controllers

import (
	"github.com/astaxie/beego"
	"github.com/PuerkitoBio/goquery"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Lishengzxc"] = "lishengzxc"
	c.TplName = "index.tpl"
}

func (c *MainController) GetList() {
	slice := make([]map[string]interface{}, 0)

	doc, _ := goquery.NewDocument("https://github.com/lishengzxc/bblog/issues")

	doc.Find(".js-navigation-item").Each(func(i int, s *goquery.Selection) {
		item := s.Find(".js-navigation-open")

		title := item.Text()
		idString, _ := item.Attr("href")

		if title != "" {
			slice = append(slice, map[string]interface{}{
				"title": title,
				"id": idString[25:],
				})
		}
  	})

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"data": slice,
	}

	c.ServeJSON()
}

func (c *MainController) GetDetail() {
	id := c.Ctx.Input.Param(":id")

	doc, _ := goquery.NewDocument("https://github.com/lishengzxc/bblog/issues/" + id)
	body, _ := doc.Find(".js-comment-container").First().Find(".js-comment-body").Html()

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"data": body,
	}

	c.ServeJSON()
}
