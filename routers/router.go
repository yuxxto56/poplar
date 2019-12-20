// @APIVersion 1.0.0
//路由配置
package routers

import (
	"poplar/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/user",&controllers.UserController{},"*:GetUser"),
		beego.NSRouter("/mem", &controllers.UserController{}, "*:Memcache"),
		beego.NSRouter("/redis", &controllers.UserController{}, "*:Redis"),
		beego.NSRouter("/rpcx", &controllers.UserController{}, "*:Rpcx"),
		beego.NSRouter("/sphinx", &controllers.UserController{}, "*:Sphinx"),
    )
	ns1 := beego.NewNamespace("/v2",
		beego.NSRouter("/user",&controllers.UserController{},"*:GetUser2"),
	)
	beego.AddNamespace(ns,ns1)
}
