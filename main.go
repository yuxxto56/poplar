package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/smallnest/rpcx/server"
	"poplar/routers"
)
//配置路由路径大小写敏感度
func configureRouterCase(){
	if ok,err := beego.AppConfig.Bool("routercase");err == nil{
		beego.BConfig.RouterCaseSensitive = ok
	}
}

//主执行函数
func main() {
	//开启协程执行rpc服务
	go runRpc()
	//配置路由路径敏感度
	configureRouterCase()
	//打印环境变量
	logs.Info("Environment Variable:MSF_ENV:",beego.BConfig.RunMode)
	//开启平滑重启
	beego.BConfig.Listen.Graceful = true
	//启动服务
	beego.Run("0.0.0.0:8000")
}

//实现rpc服务
func runRpc() {
	rpcServer := server.NewServer()
	routers.InitRpcRouters(rpcServer)
	address := fmt.Sprintf("%v:%v",beego.AppConfig.String("rpc.host"),beego.AppConfig.String("rpc.port"))
	if err := rpcServer.Serve("tcp", address); err != nil {
		//rpc启动失败
		logs.Info("failed to rpcserve:%v",err)
	}
}