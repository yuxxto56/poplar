//配置基础数据库引擎
//配置基础数据库参数
package base

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

var(
	err error
)

//初始化驱动
func setDbDriver(){
	//设置驱动数据库连接参数
	dataSource := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s",beego.AppConfig.String("db.user"),beego.AppConfig.String("db.pwd"),
		beego.AppConfig.String("db.host"),beego.AppConfig.String("db.port"),beego.AppConfig.String("db.name"),beego.AppConfig.String("db.charset"))
	maxIdle,_:= strconv.Atoi(beego.AppConfig.DefaultString("db.maxidle","10"))
	maxConn,_:= strconv.Atoi(beego.AppConfig.DefaultString("db.maxconn","0"))
	//设置注册数据库
	if err == nil{
		err = orm.RegisterDataBase("default", beego.AppConfig.String("db.type"), dataSource,maxIdle,maxConn)
	}
}