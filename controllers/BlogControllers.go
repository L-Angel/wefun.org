package controllers

import (
//	"github.com/DeanChina/wefun.org/models"
	_ "github.com/astaxie/beego"
	"time"
)

type BlogController struct {
	baseController
}

/*
存储文章（文章id，标题，发布时间，作者，内容，分类名称，分类编号）
*/
type Blog struct {
	B_Id          int
	B_Title       string
	B_ReleaseTime time.Time
	B_Author      string
	B_Content     string
	B_Type        string
	B_TypeId      int
}

func NewBlog(BlogId int, Title string, ReleaseTime time.Time, Author string, Content, Type string, TypeId int) Blog {

	return Blog{B_Id: BlogId, B_Title: Title, B_ReleaseTime: ReleaseTime, B_Author: Author, B_Content: Content, B_Type: Type, B_TypeId: TypeId}
}
func (c *BlogController) Get() {
	//BlogInfos, _ := models.GetAllViewBlogInfo(0, 10)

	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
	c.Data["IsBlog"] = true

	Blogs := make([]Blog, 2)
	Blogs[0] = NewBlog(0, "Docker 不再是单一的 Docker——记 Docker 1.11.0 版本变更", time.Now().UTC(), "wangwang", "春节过后，时隔两个月，Docker 的又一个大版本 1.11.0 发布，除了功能的不断完善之外，该版本最大的看点无疑在于 Docker Daemon 的架构由原来一个模块，现在拆分为 4 个独立的模块。", "干货", 0)
	Blogs[1] = NewBlog(1, "Docker 不再是单一的 Docker——记 Docker 1.11.0 版本变更", time.Now().UTC(), "wangwang", "春节过后，时隔两个月，Docker 的又一个大版本 1.11.0 发布，除了功能的不断完善之外，该版本最大的看点无疑在于 Docker Daemon 的架构由原来一个模块，现在拆分为 4 个独立的模块。", "干货", 0)
	c.Data["Blogs"] = Blogs
	c.TplName = "Blog.tpl"
}
