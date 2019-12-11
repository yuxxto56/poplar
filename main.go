package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "poplar/routers"
)
//配置路由路径大小写敏感度
func configureRouterCase(){
	if ok,err := beego.AppConfig.Bool("routercase");err == nil{
		beego.BConfig.RouterCaseSensitive = ok
	}
}

func main() {
	//配置路由路径敏感度
	configureRouterCase()
	//打印环境变量
	logs.Info("Environment Variable:MSF_ENV:",beego.BConfig.RunMode)
	//开启平滑重启
	beego.BConfig.Listen.Graceful = true
	//启动服务
	beego.Run("0.0.0.0:8000")

}
