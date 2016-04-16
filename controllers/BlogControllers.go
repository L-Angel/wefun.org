package controllers

import (
	_ "github.com/astaxie/beego"
)

type BlogController struct {
	baseController
}

func (c *BlogController) Get() {
	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
        c.Data["IsBlog"] = true
	c.TplName = "Blog.tpl"
}
