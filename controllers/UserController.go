package controllers

import (
     
    "github.com/astaxie/beego"
    "github.com/DeanChina/wefun.org/controllers/common"
    "fmt"
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
	fmt.Println(common.AES(password))
	fmt.Println(common.AESDecrypt(common.AES(password)))
	UserInfo := common.SetUserInfo("administrator","XXXXXX",nil)
	if username=="admin"&&password=="admin"{
	    this.Data["json"]=common.SetResponse("true",UserInfo,200)
	}else{
	    this.Data["json"]=common.SetResponse("false","password is error",nil)
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
    common.SetResponse("true","success",nil)
    common.SetResponse("email",email,nil)
    this.ServeJSON()
	return

}

/*
UserInfo:
*/

type UserInfoController struct{
	beego.Controller
}

func (this *UserInfoController) Get(){
	
	return 
}
