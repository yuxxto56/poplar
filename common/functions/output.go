package functions

import (
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego/context"
)

//正确输出
func OutApp(ctx *context.Context,params ...interface{}){
	var msg   string
	var total int
	var odata interface{}
	if len(params) > 0{
		odata = params[0]
	    if len(params) > 1 {
			if _, ok := params[1].(int); ok {
				total = params[1].(int)
			} else {
				total, _ = strconv.Atoi(params[1].(string))
			}
		}
		if len(params) > 2{
			if _,ok := params[2].(string);ok{
				msg = params[2].(string)
			}
		}
	}else{
		odata = make([]interface{},0)
	}
	maps := map[string]interface{}{
		"error":0,
		"errorMsg":msg,
		"total":total,
		"data":odata,
	}
	var content []byte
	ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	content,_= json.Marshal(maps)
	jsonS := string(content)
	if ctx.Input.Query("jsonCallback") != "" {
		jsonS = "("+jsonS+")"
	}
	content  = []byte(StringsToJSON(jsonS))
	ctx.Output.Body(content)
}


//错误输出
func ErrorApp(ctx *context.Context,errMsg string){
	odata := make([]interface{},0)
	maps := map[string]interface{}{
		"error":1,
		"errorMsg":errMsg,
		"total":false,
		"data":odata,
	}
	var content []byte
	ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	content,_= json.Marshal(maps)
	jsonS := string(content)
	if ctx.Input.Query("jsonCallback") != "" {
		jsonS = "("+jsonS+")"
	}
	content  = []byte(StringsToJSON(jsonS))
	ctx.Output.Body(content)
}
