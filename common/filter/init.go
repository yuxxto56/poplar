package filter

import (
	"github.com/astaxie/beego"
)

func init(){
	beego.InsertFilter("/*",beego.BeforeExec, runBefore )
	beego.InsertFilter("/*",beego.AfterExec, runAfter, false )
}