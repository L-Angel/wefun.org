package controllers

import (
	_ "github.com/astaxie/beego"
)

type TeamController struct {
	baseController
}

func (c *TeamController) Get() {
	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
    c.Data["IsTeam"] = true

	c.TplName = "Team.tpl"
}
