package common

func SetResponse(r interface{},m interface{},s interface{}) (map[string]interface{}){
	return map[string]interface{}{"result":r,"msg":m,"status":s}
}
func SetUserInfo(p interface{},uqid interface{},utype interface{}) (map[string]interface{}){
    return map[string]interface{}{"power":p,"uniqid":uqid,"usertype":utype}
}