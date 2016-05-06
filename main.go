package main

import (
	"github.com/DeanChina/wefun.org/controllers"
	"github.com/DeanChina/wefun.org/models"
	_ "github.com/DeanChina/wefun.org/routers"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

const (
	APP_VER = "1.0.0"
)

func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)
	beego.Info(controllers.IsPro)
    
	//connect to database
	models.RegisterDataBase()

	//demo : operate  data from database
	/*
		        o := orm.NewOrm()
		        o.Using("default")
		        user := new(models.User)
		        user.Username = "admin"
			user.Password = "admin"
			o.Insert(user)
	*/
	controllers.InitApp()
	if !controllers.IsPro {
		beego.SetStaticPath("/static_source", "static_source")
		beego.BConfig.WebConfig.DirectoryIndex = true
	}

	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
