package functions

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	lang2 "poplar/lang"
)

//正确输出
func OutApp(ctx *context.Context,params ...interface{}){
	var msg   string
	var odata interface{}
	if len(params) > 0{
		odata = params[0]
		if len(params) > 1{
			if _,ok := params[1].(string);ok{
				msg = params[1].(string)
			}
		}
	}else{
		odata = make([]interface{},0)
	}
	maps := map[string]interface{}{
		"error":0,
		"errorMsg":msg,
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
func ErrorApp(ctx *context.Context,errCode interface{}){
	odata := make([]interface{},0)
	lang := &lang2.Load{}
	errMsg := lang.GetLang(errCode.(string))
	maps := map[string]interface{}{
		"error":errCode,
		"errorMsg":errMsg,
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


