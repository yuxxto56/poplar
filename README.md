# Poplar
poplar is a http web framework written in golang
# Features
* RESTful support
* MVC architecture
* Modularity
* Annotation router
* Namespace
* Powerful development tools
* Full stack for Web & API
* Rpc && Middleware
# Quick Start
### Download and install
```cassandraql
go get -u github.com/yuxxto56/poplar
```
### Main.go
```cassandraql
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
	//启动服务
	beego.Run("0.0.0.0:8000")
}


```
### Build and Run
```cassandraql
go build main.go
./main
```

