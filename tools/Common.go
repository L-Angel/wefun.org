package tools

import(
)
func SetLoginResponse(r interface{},m interface{},s interface{}) (map[string]interface{}){
	return map[string]interface{}{"auth":r,"msg":m,"statuscode":s}
}
func SetUserInfo(username interface{},uqid interface{},role interface{},g interface{}) (map[string]interface{}){
    return map[string]interface{}{"username":username,"uniqid":uqid,"role":role,"group":g}
}
func SetResponse(r interface{},m interface{},s interface{}) (map[string]interface{}){
	return map[string]interface{}{"result":r,"msg":m,"statuscode":s}
}