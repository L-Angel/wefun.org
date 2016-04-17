package controllers

import (
	_ "github.com/astaxie/beego"
)

type MainController struct {
	//beego.Controller
	baseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
	c.Data["IsHome"] = true
	c.TplName = "index.tpl"
}
