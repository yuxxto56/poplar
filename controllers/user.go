package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"poplar/common/functions"
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
	//functions.OutApp(u.Ctx,[]map[string]string{{"id":"12121"}})
	functions.ErrorApp(u.Ctx,"字符集不能为空")
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

//@Title 测试Redis
//@router /redis [get]
func (u *UserController) Redis()  {
	var people = People{
		Name:"lilei",
		Age:18,
	}

	fmt.Println("===测试队列=========================================")
	pushData, _ := json.Marshal( people )
	err := toolLib.RedisGlobMgr.Rpush("List_test_01", string( pushData ) )
	fmt.Println( "push error:", err )
	popdata, err := redis.Bytes( toolLib.RedisGlobMgr.Lpop("List_test_01") )
	var popStruct People
	json.Unmarshal( popdata, &popStruct)
	fmt.Println("pop data:", popStruct)
	pushData1,_ := json.Marshal( People{
		Name:"hanmeimei",
		Age:16,
	})
	pushData2,_ := json.Marshal( People{
		Name:"王叔叔",
		Age:31,
	})
	toolLib.RedisGlobMgr.Rpush("List_test_01", string( pushData1 ) )
	toolLib.RedisGlobMgr.Rpush("List_test_01", string( pushData2 ) )
	lrangeData, _ := redis.ByteSlices( toolLib.RedisGlobMgr.Lrange( "List_test_01", 0, 20 ) )
	var students []People
	for _, v := range lrangeData{
		json.Unmarshal( v, &popStruct )
		students = append( students, popStruct )
	}
	fmt.Println("Lrange data:", students )
	fmt.Println(  toolLib.RedisGlobMgr.Llen("List_test_01") )


	fmt.Println("===测试Set&Get=========================================")
	type MYMAP map[string]int
	setData := MYMAP{
		"key1":1,
		"key2":2,
	}
	newSetData, _ := json.Marshal( setData )
	err = toolLib.RedisGlobMgr.Set("map_test_01", string(newSetData) )
	fmt.Println("set error :",  err)
	var getData MYMAP
	gdata , _ := redis.Bytes( toolLib.RedisGlobMgr.Get( "map_test_01" ) )
	json.Unmarshal( gdata, &getData )
	fmt.Println( "get data:", getData )

	toolLib.RedisGlobMgr.Set("incrtest", 10 )
	toolLib.RedisGlobMgr.Incrnum("incrtest", 3)
	inc, _ := redis.Int( toolLib.RedisGlobMgr.Get("incrtest") )
	fmt.Println("获取加3后的值：", inc )

	fmt.Println("===测试集合=========================================")
	sData := People{
		Name:"吴彦祖",
		Age:35,
	}
	newSdata, _ := json.Marshal( sData )
	toolLib.RedisGlobMgr.Sadd("skey_test_01", string( newSdata ) )
	fmt.Print("检查值是否存在：")
	fmt.Println( toolLib.RedisGlobMgr.Scontains( "skey_test_01", string(newSdata) ) )

	sData2 := People{
		Name:"刘德华",
		Age:55,
	}
	newSdatas, _ := json.Marshal( sData2 )
	toolLib.RedisGlobMgr.Sadd("skey_test_01", string( newSdatas ) )

	//获取结合里面的所有数据
	byteDatas, err := redis.ByteSlices( toolLib.RedisGlobMgr.Smembers("skey_test_01") )
	if  err != nil {
		fmt.Println(err)
	}
	var out People
	var out4 []People
	for _, v := range byteDatas{
		json.Unmarshal( v, &out )
		out4 = append( out4, out )
	}
	fmt.Println( "取出集合里面所有数据：", out4 )
	//移除一个数据
	fmt.Print( toolLib.RedisGlobMgr.Srem("skey_test_01", string( newSdata ) ) )
	byteDatas, err = redis.ByteSlices( toolLib.RedisGlobMgr.Smembers("skey_test_01") )
	if  err != nil {
		fmt.Println(err)
	}

	var out5 []People
	for _, v := range byteDatas{
		json.Unmarshal( v, &out )
		out5 = append( out5, out )
	}
	fmt.Println( "重新取出集合里面所有数据：", out5 )



	fmt.Println("===测试HASH=========================================")

	var hashData int = 10
	fmt.Println( toolLib.RedisGlobMgr.Hset("hashkey_test_01", "key001", hashData ) )
	fmt.Print("获取key001:")
	fmt.Println( redis.Int( toolLib.RedisGlobMgr.Hget( "hashkey_test_01", "key001" ) ) )
	fmt.Println( toolLib.RedisGlobMgr.Hincrby("hashkey_test_01", "key001", 6 ) )
	fmt.Print("加6后，重新获取key001:")
	fmt.Println( redis.Int( toolLib.RedisGlobMgr.Hget( "hashkey_test_01", "key001" ) ) )


	fmt.Println("===测试有序集合=========================================")
	toolLib.RedisGlobMgr.Zset("zkey", 1.1, "xiaowang" )
	toolLib.RedisGlobMgr.Zset("zkey", 1, "xiaoHong" )
	toolLib.RedisGlobMgr.Zset("zkey", 1.6, "xiaoMing" )
	toolLib.RedisGlobMgr.Zset("zkey", 1.2, "xiaoLan" )

	zdata, _ := toolLib.RedisGlobMgr.Zrange("zkey", 0, 3, "asc", true )
	fmt.Print("按照升序获取有序集合里面的值，并且返回score的值：")
	fmt.Println( redis.StringMap( zdata, nil ))

	fmt.Print("按照降序获取有序集合里面的值，不返回score的值：")
	fmt.Println(redis.Strings( toolLib.RedisGlobMgr.Zrange("zkey", 0, 3, "desc", false ) ) )

	fmt.Println("删除xiaoLan:")
	toolLib.RedisGlobMgr.Zdel("zkey", "xiaoLan")
	fmt.Print("重新按照降序获取有序集合里面的值，不返回score的值：")
	fmt.Println(redis.Strings( toolLib.RedisGlobMgr.Zrange("zkey", 0, 3, "asc", false ) ) )
	fmt.Print("当前集合长度")
	fmt.Println(toolLib.RedisGlobMgr.Zcount("zkey", 0, 100 ) )
	fmt.Print("获取成员xiaoMing的排名：")
	fmt.Println(toolLib.RedisGlobMgr.Zrank("zkey", "xiaoMing", "asc" ) )


	fmt.Println("===测试切换db=========================================")
	var redisDb2 = &toolLib.RedisMgr{}
	toolLib.RedisInit(2, redisDb2 )
	redisDb2.Set("keydb2", 111 )
	redisDb2.Set("keydb2_1", 112 )
	fmt.Print("测试切换db, 从db2拿数据：")
	fmt.Println( redis.Int( redisDb2.Get("keydb2") )  )
	//释放掉redisDb2
	redisDb2.Close()
	redisDb2 = nil

	u.Ctx.WriteString("end")
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
	u.Ctx.WriteString("end")
}

func (u *UserController) Rpcx()  {
	/*args := &_interface.Reply{}
	reply := &[]map[string]interface{}{}

	rpcStudent := new(poplar.Student).Init()
	defer rpcStudent.Destruct()

	if err := rpcStudent.GetAll(context.Background(), args, reply); err!=nil{
		log.Fatalf("failed to call: %v", err)
	}


	u.Data["json"] = reply*/
	u.ServeJSON()
}
