package controllers

import (
	"github.com/DeanChina/wefun.org/models"
	_ "github.com/astaxie/beego"
)

type BlogController struct {
	baseController
}

func (c *BlogController) Get() {
	BlogInfos, _ := models.GetAllViewBlogInfo(0, 10)

	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
	c.Data["IsBlog"] = true
	c.Data["BlogInfos"] = BlogInfos
	c.TplName = "Blog.tpl"
}
