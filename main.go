package main

import (
	_ "wefun.org/routers"
	"wefun.org/controllers"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

const (
        APP_VER = "1.0.0"
)

func main() {
        beego.Info(beego.BConfig.AppName, APP_VER)
        beego.Info(controllers.IsPro)

        controllers.InitApp()
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}

