package toolLib

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bradfitz/gomemcache/memcache"
	"poplar/common/functions"
	"reflect"
	"time"
)

const (
	// 读写超时时间 单位毫秒
	MemcacheTimeout = 600

	//最大空闲连接数
	MemcacheMaxIdleConns = 30
)

var MemMgr *MemcacheMgr

type MemcacheMgr struct {
	Host string
	Port int
	Client *memcache.Client
}

func init()  {
	port, _ := beego.AppConfig.Int("memcache.port")
	MemMgr = &MemcacheMgr{
		Host: beego.AppConfig.String("memcache.host"),
		Port: port,
	}
	MemMgr.Client = memcache.New( fmt.Sprintf( "%s:%d", MemMgr.Host, MemMgr.Port ) )
	if MemMgr.Client == nil{
		fmt.Println("Memcache 创建失败")
	}
	//设置读写超时时间
	timeOut, err := beego.AppConfig.Int64("memcache.timeOut")
	if err  != nil{
		timeOut = int64( MemcacheTimeout )
	}
	MemMgr.Client.Timeout =  time.Duration( timeOut )  * time.Millisecond
	//设置最大空闲连接数
	maxIdleConns, err := beego.AppConfig.Int("memcache.maxIdleConns")
	if err != nil{
		maxIdleConns = MemcacheMaxIdleConns
	}
	MemMgr.Client.MaxIdleConns = maxIdleConns
}


//获取以Gob格式编码存储的数据
func ( m *MemcacheMgr ) GetGob( key string, toData interface{} ) error {
	item, err := m.Client.Get( key )
	if err != nil{
		return  err
	}
	err = functions.GobDecodeByte( item.Value, toData )
	if err != nil{
		return err
	}
	return  nil
}

//要清除的缓存
func ( m *MemcacheMgr )  Clear( key string ) error {
	return  m.Client.Delete( key )
}

//生成缓存 - 数据以Gob编码格式存储
func ( m *MemcacheMgr ) SetGob( key string, data interface{}, expire int32 ) error {
	fmt.Println( reflect.TypeOf(data))
	byteData, err :=  functions.GobEncode2Byte( data )
	if err != nil{
		return  err
	}
	return m.Client.Set( &memcache.Item{
		Key:key,
		Value:byteData,
		Expiration: expire,
	})
}

//添加缓存 - 数据以Gob编码格式存储
func ( m *MemcacheMgr ) AddGob(key string, data interface{}, expire int32 ) error {
	byteData, err :=  functions.GobEncode2Byte( data )
	if err != nil{
		return  err
	}
	return  m.Client.Add( &memcache.Item{
		Key:key,
		Value:byteData,
		Expiration: expire,
	})
}

//正常形式的添加缓存
func (m *MemcacheMgr) Set(key string, data []byte, expire int32 ) error {
	return m.Client.Set( &memcache.Item{
		Key:key,
		Value:data,
		Expiration: expire,
	})
}

//正常形式的添加缓存
func (m *MemcacheMgr) Add(key string, data []byte, expire int32 ) error {
	return m.Client.Add( &memcache.Item{
		Key:key,
		Value:data,
		Expiration: expire,
	})
}

//正常形式的获取数据
func (m *MemcacheMgr) Get(key string ) ([]byte, error) {
	item, err := m.Client.Get(  key )
	if err != nil{
		return nil, err
	}
	return item.Value,nil
}

//加法，只能对数值型缓存使用
func ( m *MemcacheMgr ) Increment( key string, value uint64 ) ( uint64, error ){
	return m.Client.Increment( key, value )
}

//减法，只能对数值型缓存使用
func ( m *MemcacheMgr ) Decrement( key string, value uint64 ) ( uint64, error ) {
	return m.Client.Decrement( key, value )
}



