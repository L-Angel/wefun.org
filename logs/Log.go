package log

import (
    "github.com/astaxie/beego/logs"
)


func InitLog(){
	log:=logs.NewLogger(10000)
    log.SetLogger("file",`{"filename":"wefun.org.log"}`)
    log.Debug("this is a debug message")
}
