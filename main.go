package main

import (
	"github.com/astaxie/beego/logs"
	"os"
	_ "poplar/routers"
	"strings"

	"github.com/astaxie/beego"
)
const (
	envName = "MSF_ENV"
)
//配置环境
func configureEnv(){
	runMode := os.Getenv(envName)
	if runMode == ""{
		runMode = "dev"
	}else{
		runMode = strings.ToLower(runMode)
	}
	beego.BConfig.RunMode = runMode
}
//配置路由路径大小写敏感度
func configRouterCase(){
	if ok,err := beego.AppConfig.Bool("routercase");err == nil{
		beego.BConfig.RouterCaseSensitive = ok
	}
}

func main() {
	//配置环境
	configureEnv()
	//配置路由路径敏感度
	configRouterCase()

	//打印环境变量
	logs.Info("环境变量：",beego.BConfig.RunMode)
	//启动服务
	beego.Run("0.0.0.0:8000")
}
