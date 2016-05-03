package controllers

import (
     
    "github.com/astaxie/beego"
    "github.com/DeanChina/wefun.org/controllers/common"
)
//notice : modify method get to post when in Produce environment
/**
login:
Response value :
result: true | false | error
**/
type LoginController struct{
	beego.Controller
}

func (this *LoginController) Get() {
	username:=this.GetString("username")
	password:=this.GetString("password")
	if username=="admin"&&password=="admin"{
	    this.Data["json"]=common.SetResponse("true","success")
	}else{
	    this.Data["json"]=common.SetResponse("false","password is error")
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
type RegisterController struct{
	beego.Controller
}
func (this *RegisterController) Get(){
	//username:=this.GetString("username")
	//password:=this.GetString("password")
	email:=this.GetString("email")
	//tel:=this.GetString("tel")
	//address:=this.GetString("address")
    common.SetResponse("true","success")
    common.SetResponse("email",email)
    this.ServeJSON()
	return

}