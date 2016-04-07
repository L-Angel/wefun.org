package routers

import (
	"github.com/DeanChina/wefun.org/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
