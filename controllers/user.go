package controllers

import (
	//"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"poplar/common/models/base"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (u *UserController) GetUser(){



	/*dataSource := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s",beego.AppConfig.String("db::user"),beego.AppConfig.String("db::pwd"),
		beego.AppConfig.String("db::host"),beego.AppConfig.String("db::port"),beego.AppConfig.String("db::name"),beego.AppConfig.String("db::charset"))
	logs.Info("dataSourceUser:",dataSource)*/

	m := base.NewModel("student")
	maps := map[string]interface{}{"id":[]string{"in","1,2,3"}}

	result := m.Where(maps).Find()

	//u.Ctx.WriteString(str)
	u.Data["json"] = result
    u.ServeJSON()

}

func (u *UserController) GetUser2(){
	//u.Ctx.WriteString("getUser")
	u.Data["json"] = map[string]string{"user2":"liyang2"}
	u.ServeJSON()
}


// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

