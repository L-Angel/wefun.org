package controllers

import (
	_ "github.com/astaxie/beego"
)

type LifeController struct {
	baseController
}

func (c *LifeController) Get() {
	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
        c.Data["IsLife"] = true
	c.TplName = "Life.tpl"
}
