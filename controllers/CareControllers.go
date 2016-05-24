package controllers

import (
	_ "github.com/astaxie/beego"
)

type CareController struct {
	baseController
}

func (c *CareController) Get() {
	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
	c.Data["IsCare"] = true
	c.TplName = "Care.tpl"
}
