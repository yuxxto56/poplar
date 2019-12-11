package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"poplar/common/models"
	"poplar/common/toolLib"
	"strconv"
	"time"
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
	//data := map[string]interface{}{"name":"lidayang","age":22}
	sdu := new(models.StudentModel).Init()
	//result := sdu.Insert(data)
	result := sdu.GetAll()
	bytes,_ := json.Marshal(result)

	//fmt.Println("result:",result)
	fmt.Println("lastSql:",sdu.Model.GetLastSql())
	fmt.Println(fmt.Sprintf("timer:%s,result:%s",time.Now(),string(bytes[:])))
	u.Data["json"] = result
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

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

