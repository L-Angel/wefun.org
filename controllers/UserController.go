package controllers

import (
    "github.com/DeanChina/wefun.org/models"
    "github.com/DeanChina/wefun.org/tools"
    "github.com/astaxie/beego/cache"
    "github.com/astaxie/beego/utils/captcha"
    "strings"
    "fmt"
)

var cpt *captcha.Captcha

func init(){
	store := cache.NewMemoryCache()
	cpt=captcha.NewWithFilter("/captcha/",store)
}

type LoginController struct{
	baseController
}
func (this *LoginController) Post(){
	this.Data["Website"] = "www.wefun.org"
	this.Data["Email"] = "Dean.Wefun@gmail.com"
    this.Data["IsLogin"] = true
    this.Data["Success"] =cpt.VerifyReq(this.Ctx.Request)
	this.TplName="login.tpl"
}

func (this *LoginController) Get(){
	this.Data["Website"] = "www.wefun.org"
	this.Data["Email"] = "Dean.Wefun@gmail.com"
    this.Data["IsLogin"] = true
	this.TplName="login.tpl"
}



type RegisterController struct{
	baseController
}
func (this *RegisterController) Post(){
	this.Data["Website"] = "www.wefun.org"
	this.Data["Email"] = "Dean.Wefun@gmail.com"
    this.Data["IsRegister"] = true
    this.Data["Success"]=cpt.VerifyReq(this.Ctx.Request)
    this.TplName="register.tpl"
}

func (this *RegisterController) Get(){
	this.Data["Website"] = "www.wefun.org"
	this.Data["Email"] = "Dean.Wefun@gmail.com"
    this.Data["IsRegister"] = true
    this.TplName = "register.tpl"
}
//notice : modify method get to post when in Produce environment
/**
login:
Response value :
result: true | false | error
**/
type InterfaceLoginController struct{
	baseController
}

func (this *InterfaceLoginController) Get() {
	username:=strings.TrimSpace(this.GetString("username"))
	password:=strings.TrimSpace(this.GetString("password"))
	captcha_id:=this.GetString("captcha_id")
	captcha:=this.GetString("captcha")
	b:=cpt.Verify(captcha_id,captcha)
	//fmt.Println(tools.AES(password))
	//fmt.Println(tools.AESDecrypt(tools.AES(password)))
	if b!=true{
		this.Data["json"]=tools.SetLoginResponse("failure","Captcha Verify not match.",20010)
	}else if username != "" {
		user,err := models.GetUserByUsername(username)
		fmt.Println(user)
		fmt.Println(err)
		if user != nil && user.Password != "" && user.Password == tools.AES(password) {
			UserInfo := tools.SetUserInfo(user.Username,"uniqid000000001","administrator","administrator")
			this.Data["json"]=tools.SetLoginResponse("success",UserInfo,10000)
		}else if user != nil && user.Password != tools.AES(password) {
            this.Data["json"]=tools.SetLoginResponse("failure","Password is error.",20001)
		}else if user == nil {
		    this.Data["json"]=tools.SetLoginResponse("failure","dont find user "+username,20002)	
		}
	}else{
	    this.Data["json"]=tools.SetLoginResponse("error","username is nil",30000)
	}

	this.ServeJSON()
	return 
}

/**
Register User:
username
password
email
tel
address 
**/
type InterfaceRegisterController struct{
	baseController
}

func (this *InterfaceRegisterController) Get(){
	username:=this.GetString("username")
	password:=tools.AES(this.GetString("password"))
	email:=this.GetString("email")
	tel:=this.GetString("tel")
	address:=this.GetString("address")
	captcha_id:=this.GetString("captcha_id")
	captcha:=this.GetString("captcha")
	b:=cpt.Verify(captcha_id,captcha)

	if b != true {
		this.Data["json"]=tools.SetResponse("failure","Captcha Verify not match",20010)
	}else{
		result,msg,err:=models.CheckOrAddUser(username,password,email,tel,address)
		if result == "true"{
    			this.Data["json"]=tools.SetResponse("success",msg,10000)
		}else if result == "false"{
			if msg=="existed" {
				this.Data["json"]=tools.SetResponse("failure",msg,20001)
			}else{
				this.Data["json"]=tools.SetResponse("failure",msg,20000)
			}
		}else if result == "error"{
    		        this.Data["json"]=tools.SetResponse("error",err,30000)
		}else{
			this.Data["json"]=tools.SetResponse("error","dont clear error",30010)
		}

         }
        this.ServeJSON()
	return
}

/*
UserInfo:
*/

type UserInfoController struct{
	baseController
}

func (this *UserInfoController) Get(){
	
	return 
}
