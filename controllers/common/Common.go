package common

func SetResponse(r interface{},m interface{})(map[string]interface{}){
	return map[string]interface{}{"result":r,"msg":m}
}