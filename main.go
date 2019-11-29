package main

import (
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

func main() {
	//配置环境
	configureEnv()
	//启动服务
	beego.Run()
}
