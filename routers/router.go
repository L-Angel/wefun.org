package routers

import (
	"github.com/DeanChina/wefun.org/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.ErrorController(&controllers.ErrorController{})
    beego.Router("/", &controllers.MainController{})
    beego.Router("/blog", &controllers.BlogController{})
    beego.Router("/community", &controllers.CommunityController{})
    beego.Router("/docs", &controllers.DocsController{})
    beego.Router("/life", &controllers.LifeController{})
    beego.Router("/care", &controllers.CareController{})
    beego.Router("/join", &controllers.JoinController{})
    beego.Router("/about", &controllers.AboutController{})
    beego.Router("/team", &controllers.TeamController{})
    beego.Router("/user/login",&controllers.LoginController{})
    beego.Router("/user/register",&controllers.RegisterController{})
    beego.Router("/md/*.md",   &controllers.MarkdownController{})
}
