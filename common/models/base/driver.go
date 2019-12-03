//配置基础数据库引擎
//配置基础数据库参数
package base

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//初始化驱动
func init(){
	//选择数据驱动
	_:= orm.RegisterDriver(beego.AppConfig.String("db::type"), orm.DRMySQL)
	//设置数据库参数
	dataSource := fmt.Sprintf("%s:%s@%s/%s?charset=%s",beego.AppConfig.String("db::user"),beego.AppConfig.String("db::pwd"),
		beego.AppConfig.String("db::host"),beego.AppConfig.String("db::name"),beego.AppConfig.String("db::charset"))
	if err:=orm.RegisterDataBase("default", beego.AppConfig.String("db::type"), dataSource);err !=nil{
		logs.Error("RegisterDataBase Error,Error is ",err.Error());
		return
	}
}