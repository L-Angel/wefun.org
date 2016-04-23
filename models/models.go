package models

import (
	//"fmt"
	//"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterDataBase() {
	var DbUser = beego.AppConfig.String("dbuser")
	var DbPass = beego.AppConfig.String("dbpass")
	var DbHosts = beego.AppConfig.String("dbhosts")
	//var DbType = beego.AppConfig.String("dbtype")
	var DbName = beego.AppConfig.String("dbname")

	var ConnString = DbUser + ":" + DbPass + "@tcp(" + DbHosts + ")/" + DbName
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", ConnString)

}
