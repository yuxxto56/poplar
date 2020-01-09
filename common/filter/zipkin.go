package filter

import (
	"fmt"
	"github.com/astaxie/beego/context"
	rpcxplugins "github.com/nosixtools/rpcx-plugins/opentracing"
	"github.com/opentracing/opentracing-go"
)

var runBefore = func (ctx *context.Context){
	fmt.Println("run before")
	span, cont, err := rpcxplugins.GenerateSpanWithContext( ctx.Request.Context(), fmt.Sprintf("请求地址：%s", ctx.Request.RequestURI  ))
	ctx.Request = ctx.Request.WithContext( cont )
	if err != nil{
		fmt.Println( "创建span错误", err )
	}
	ctx.Input.SetData("zipkinSpan", span )
}

var runAfter = func( ctx *context.Context ) {
	fmt.Println("run after")
	span := ctx.Input.GetData("zipkinSpan" )
	if span, ok := span.(opentracing.Span); ok == true {
		defer span.Finish()
		span.LogKV("结束")
	}
}

