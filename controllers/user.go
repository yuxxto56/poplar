package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	models2 "poplar/common/models"
	"poplar/common/models/base"
	"poplar/common/toolLib"
	"strconv"

)

// Operations about Users
type UserController struct {
	beego.Controller
}

type People struct {
	Name string
	Age  int
}

func (u *UserController) GetUser(){

	m := base.NewModel("student")
	//maps := map[string]interface{}{"id":1}
	//res,_:= m.Where(maps).Data(map[string]interface{}{"name":"liii","age":38}).Update()
	//map
	res := m.Where(map[string]interface{}{"id":"1"}).Count("id")
	fmt.Println("res",res)
	u.ServeJSON()
}

func (u *UserController) GetUser2(){
	//u.Ctx.WriteString("getUser")
	u.Data["json"] = map[string]string{"user2":"liyang2"}
	u.ServeJSON()
}

func ( u *UserController ) Memcache()  {
	var (
		structKey string = "test01"
		mapKey string = "test02"
		strKey string = "test03"
		intKey string = "test04"
		incrKey string= "test05"
	)

	//结构体
	var inmdata = People{
		Name:"lilei",
		Age:18,
	}

	toolLib.MemMgr.SetGob(structKey, inmdata, 3600 )
	var out People
	toolLib.MemMgr.GetGob( structKey, &out )
	fmt.Println( "测试结构体：", out )


	var mapdata map[string]string
	mapdata = make(map[string]string)
	mapdata["name"] = "xiaohua"
	mapdata["age"] = "20"
	toolLib.MemMgr.SetGob(mapKey, mapdata, 3600 )
	var outMapdata map[string]string
	toolLib.MemMgr.GetGob( mapKey, &outMapdata )
	fmt.Println( "测试Map:", outMapdata )


	toolLib.MemMgr.SetGob(strKey, "hello baby��", 3600)
	var outStrData string
	toolLib.MemMgr.GetGob(strKey, &outStrData)
	fmt.Println("测试字符串:", outStrData)


	var intData uint64 = 16
	toolLib.MemMgr.SetGob(intKey, intData, 3600)
	var outIntData uint64
	toolLib.MemMgr.GetGob(intKey, &outIntData)
	fmt.Println("测试整数：", outIntData)


	var incrData uint64 = 1
	var outIncrData uint64
	var incrYdata string = "20"

	toolLib.MemMgr.Set(incrKey, []byte(incrYdata), 3600 )
	toolLib.MemMgr.Increment( incrKey, incrData )

	byteOut, err  := toolLib.MemMgr.Get( incrKey )
	if  err != nil{
		fmt.Println( err )
	}
	outIncrData, _ = strconv.ParseUint(string(byteOut), 10, 64)
	fmt.Println("测试递增:", outIncrData )

	u.Ctx.WriteString("end")
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models2.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models2.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models2.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models2.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models2.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models2.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models2.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models2.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
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

