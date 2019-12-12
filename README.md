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

# Create Model
* 查看生成model帮助
```cassandraql
cd tests
go run create_model.go -help
  -dev string #指定开发环境生成model,可选有 dev、prod、test 默认环境:dev
        usage config environment (default "dev")
  -pk string #指定项目包名 默认包名为:poplar
        usage package name (default "poplar")
  -table string #指定表名称 不可缺少
        usage mysql table name
```
*生成model
```cassandraql
go run create_model.go -table=student
```

# Project Structure
```cassandraql
common #公共文件包
|  |-functions #定义全局函数、方法体
|  |-models #Model层
|  |  |-base
|  |     |-driver.go #注册数据库引擎
|  |     |-model.go  #基Model
|  |-logics #Logics业务逻辑层
|  |-toolLib #第三方类库
|  |  |-memcaheMgr.go  #操作memcache库
|  |  |-RedisMgr.go    #操作redis库
conf #配置包
|  |-app.conf #公共配置
|  |-dev  #开发环境配置
|  |-prod #生产环境配置
|  |-test #测试环境配置
tests
|  |-create_model.go #生成model工具
controllers #controller层
routers #路由包
main.go #入口文件
```