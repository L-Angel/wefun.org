package logs

import (
    "github.com/astaxie/beego/logs"
)


func InitLog()(*logs.BeeLogger){
	log:=logs.NewLogger(1000)
    log.SetLogger("file",`{"filename":"wefun.org.log"}`)
    return log
}

func ErrorLog(flag string, msg error){
	log := InitLog()
	if "" != flag && nil!= msg{
		log.Error(flag+" Error : "+msg.Error())
	} else if nil != msg{
		log.Error(msg.Error())
	}
}