package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "Tom-ly.me"
	c.Data["Email"] = "3419572132@gmail.com"
	c.TplName = "index.tpl"
}
