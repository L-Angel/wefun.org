package controllers

import (
	_ "github.com/astaxie/beego"
)

type AboutController struct {
	baseController
}

func (c *AboutController) Get() {
	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
	c.Data["IsAbout"] = true
	c.TplName = "About.tpl"
}
