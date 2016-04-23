package controllers

import (
	"fmt"
	"github.com/DeanChina/wefun.org/models"
	_ "github.com/astaxie/beego"
)

type BlogController struct {
	baseController
}

func (c *BlogController) Get() {
	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
	c.Data["IsBlog"] = true
	BlogInfos, _ := models.GetAllViewBlogInfo(0, 10)
	fmt.Println(BlogInfos)
	c.Data["BlogInfos"] = BlogInfos
	c.TplName = "Blog.tpl"
}
