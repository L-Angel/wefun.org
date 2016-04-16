package controllers

import (
	_ "github.com/astaxie/beego"
)

type CommunityController struct {
	baseController
}

func (c *CommunityController) Get() {
	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
        c.Data["IsCommunity"] = true
	c.TplName = "Community.tpl"
}
