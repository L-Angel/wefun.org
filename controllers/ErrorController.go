package controllers


import(
    "github.com/astaxie/beego"
    "net/http"
)

type ErrorController struct{
	beego.Controller
}

func (this *ErrorController) Error404() {
    this.Ctx.Redirect(302,"http://www.baidu.com")
    this.EnableRender=false

}

func (this *ErrorController) Error500() {
    this.Data["content"] = "internal server error"
    this.TplName = "error/500.tpl"
}

func (this *ErrorController) ErrorDb() {
    this.Data["content"] = "database is now down"
    this.TplName = "error/errordb.tpl"
}

func PageNotFound(rw http.ResponseWriter,r *http.Request){
	rw.Header().Set("Location","www.baidu.com")
}