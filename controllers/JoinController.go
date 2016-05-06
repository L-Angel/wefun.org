package controllers

import (
	_ "github.com/astaxie/beego"
)

type JoinController struct {
	baseController
}

func (c *JoinController) Get() {
	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
    c.Data["IsJoin"] = true
	c.TplName = "Join.tpl"
}
