package routers

import (
	"github.com/DeanChina/wefun.org/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/blog", &controllers.BlogController{})
    beego.Router("/community", &controllers.CommunityController{})
    beego.Router("/docs", &controllers.DocsController{})
    beego.Router("/life", &controllers.LifeController{})
    beego.Router("/care", &controllers.CareController{})
    beego.Router("/join", &controllers.JoinController{})
    beego.Router("/about", &controllers.AboutController{})
    beego.Router("/team", &controllers.TeamController{})
}
