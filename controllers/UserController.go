package controllers

import (
    "github.com/DeanChina/wefun.org/models"
    "github.com/DeanChina/wefun.org/tools"
    "strings"
    "fmt"
)

type LoginController struct{
	baseController
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
func (this *RegisterController) Get(){
	this.Data["Website"] = "www.wefun.org"
	this.Data["Email"] = "Dean.Wefun@gmail.com"
    this.Data["IsRegister"] = true
    this.TplName="register.tpl"
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
	//fmt.Println(tools.AES(password))
	//fmt.Println(tools.AESDecrypt(tools.AES(password)))
	if username != "" {
		user,err := models.GetUserByUsername(username)
		fmt.Println(user)
		fmt.Println(err)
		if user != nil && user.Password != "" && user.Password == tools.AES(password) {
			UserInfo := tools.SetUserInfo(user.Username,"uniqid000000001","administrator","administrator")
			this.Data["json"]=tools.SetLoginResponse("success",UserInfo,10000)
		}else if user != nil && user.Password != tools.AES(password) {
            this.Data["json"]=tools.LoginSetResponse("failure","password is error",20001)
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

	result,msg,err:=models.CheckOrAddUser(username,password,email,tel,address)
	if result == "true"{
    this.Data["json"]=tools.SetResponse(result,msg,nil)
	}else if result == "false"{
    this.Data["json"]=tools.SetResponse(result,msg,nil)
	}else if result == "error"{
    this.Data["json"]=tools.SetResponse(result,err,nil)
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
