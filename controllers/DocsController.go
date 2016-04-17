package controllers

import (
	_ "github.com/astaxie/beego"
)

type DocsController struct {
	baseController
}

func (c *DocsController) Get() {
	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
	c.Data["IsDocs"] = true
	c.TplName = "Docs.tpl"
}
