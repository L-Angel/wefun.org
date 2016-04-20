package controllers

import (
	_ "github.com/astaxie/beego"
)

type MarkdownController struct {
	baseController
}

func (c *MarkdownController) Get() {
	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
        c.Data["IsMarkdown"] = true
	c.TplName = "Markdown.tpl"
}
