package toolLib

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"strings"
	"time"
)

var (
	RedisGlobMgr 		*RedisMgr = &RedisMgr{}
)


type RedisMgr struct {
	pool        *redis.Pool

}


func init()  {
	db , err := beego.AppConfig.Int("redis.db")
	if err != nil {
		db  = 0
	}
	RedisInit( db, RedisGlobMgr )
}

func RedisInit( db int, redism *RedisMgr )  {
	host := fmt.Sprintf("%s:%s", beego.AppConfig.String("redis.host"), beego.AppConfig.String("redis.port") )

	//最大空闲连接数
	maxIdle, err := beego.AppConfig.Int("redis.maxIdle")
	if err != nil{
		maxIdle = 50
	}
	//最大连接数
	maxActive, err := beego.AppConfig.Int("redis.maxActive")
	if err != nil{
		maxActive = 600
	}
	//空闲链接超时时间
	idleTimeout, err := beego.AppConfig.Int64("redis.idleTimeout")
	if err != nil{
		idleTimeout = 600
	}
	//如果超过最大连接，是报错，还是等待
	wait, err := beego.AppConfig.Bool("redis.wait")
	if err != nil{
		wait = true
	}
	timeout := time.Duration( idleTimeout )  * time.Second
	//password
	pass := beego.AppConfig.String("redis.pass")

	redism.pool = &redis.Pool{
		MaxIdle:maxIdle,
		MaxActive:maxActive,
		IdleTimeout: timeout,
		Wait:wait,
		Dial: func() (conn redis.Conn, e error) {
			con, err := redis.Dial("tcp", host,
				redis.DialPassword(pass),
				redis.DialDatabase(db),
				redis.DialConnectTimeout(2*time.Second),
				redis.DialReadTimeout(2*time.Second),
				redis.DialWriteTimeout(3*time.Second), )
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}

}


func (r *RedisMgr) GetConn() redis.Conn {
	return  r.pool.Get()
}


//向一个key[队列]的尾部添加一个元素
func (r *RedisMgr) Rpush( key string, data interface{} ) error {
	conn := r.GetConn()
	if conn.Err() != nil {
		return conn.Err()
	}

	defer func() {
		conn.Flush()
		conn.Close()

	}()
	return  conn.Send("RPUSH", key, data )
}

//向一个key[队列]的头部添加一个元素
func (r *RedisMgr) Lpush( key string, data interface{} ) error {
	conn := r.GetConn()
	if conn.Err() != nil {
		return conn.Err()
	}

	defer func() {
		conn.Flush()
		conn.Close()
	}()
	return  conn.Send("LPUSH", key, data )
}

//取出队列中第一个key取元素值
func (r *RedisMgr) Lpop( key string ) ( interface{}, error )   {
	conn := r.GetConn()
	if conn.Err() != nil {
		return nil, conn.Err()
	}
	defer func() {
		conn.Close()
	}()
	return  conn.Do("LPOP", key )
}

//返回名称为key的list中start至end之间的元素（end为 -1 ，返回所有）
func ( r *RedisMgr ) Lrange(key string, start int, end int )  ( interface{}, error )  {
	conn := r.GetConn()
	if conn.Err() != nil {
		return nil, conn.Err()
	}

	defer func() {
		conn.Close()
	}()
	return  conn.Do("LRANGE", key, start, end )
}

//获取队列长度
func ( r *RedisMgr ) Llen( key string ) (int , error) {
	conn := r.GetConn()
	if conn.Err() != nil {
		return 0, conn.Err()
	}
	defer func() {
		conn.Close()
	}()

	len, err := redis.Int( conn.Do("LLEN", key) )
	if err != nil{
		return 0,err
	}
	return len, nil
}

// 判断一个key集合里是否存在某个value值，存在返回True
func ( r *RedisMgr ) Scontains(key string, data interface{}) (bool, error)  {
	conn := r.GetConn()
	if conn.Err() != nil{
		return false,conn.Err()
	}
	defer func() {
		conn.Close()
	}()

	return  redis.Bool( conn.Do("SISMEMBER", key, data ) )
}

//向集合添加元素
func ( r *RedisMgr) Sadd(key string, data interface{})  error {
	conn := r.GetConn()
	if conn.Err() != nil{
		return conn.Err()
	}

	defer func() {
		conn.Flush()
		conn.Close()
	}()

	return conn.Send("SADD", key, data )
}


//返回key集合所有的元素
func ( r *RedisMgr )  Smembers(key string) (interface{}, error){
	conn := r.GetConn()
	if conn.Err() != nil{
		return nil, conn.Err()
	}

	defer func() {
		conn.Close()
	}()
	return  conn.Do("SMEMBERS", key )
}

//在key集合中移除指定的元素
func ( r *RedisMgr ) Srem( key string, data interface{} ) error {
	conn := r.GetConn()
	if conn.Err() != nil{
		return conn.Err()
	}

	defer func() {
		conn.Flush()
		conn.Close()
	}()

	return  conn.Send("SREM", key, data)
}

//删除指定的key
func ( r *RedisMgr ) Clear( key string) error  {
	conn := r.GetConn()
	if conn.Err() != nil{
		return conn.Err()
	}

	defer func() {
		conn.Flush()
		conn.Close()
	}()

	return conn.Send("DEL", key )
}

//设置数据
func ( r *RedisMgr) Set(key string, data interface{} ) error  {
	conn := r.GetConn()
	if conn.Err() != nil{
		return conn.Err()
	}

	defer func() {
		conn.Flush()
		conn.Close()
	}()

	return  conn.Send("SET", key, data )
}

//获取数据
func ( r *RedisMgr) Get(key string ) (interface{}, error)  {
	conn := r.GetConn()
	if conn.Err() != nil{
		return nil, conn.Err()
	}

	defer func() {
		conn.Close()
	}()


	return  conn.Do("GET", key )

}

//设置某个hashKey名称的下的keyvalue值
func ( r *RedisMgr ) Hset( hashKey string, key string, data interface{} ) error {
	conn := r.GetConn()
	if conn.Err() != nil{
		return conn.Err()
	}

	defer func() {
		conn.Flush()
		conn.Close()
	}()

	return conn.Send("HSET", hashKey, key, data )

}

//得到某个hashKey名称下的key信息
func (r *RedisMgr ) Hget( hashKey string, key string ) ( interface{}, error )  {
	conn := r.GetConn()
	if conn.Err() != nil{
		return nil, conn.Err()
	}

	defer func() {
		conn.Close()
	}()

	return conn.Do("HGET", hashKey, key )
}

//删除haskKey下面的key建
func ( r *RedisMgr) Hdel( hashKey string, key string ) error  {
	conn := r.GetConn()
	if conn.Err() != nil{
		return conn.Err()
	}

	defer func() {
		conn.Flush()
		conn.Close()
	}()

	return  conn.Send("HDEL", hashKey, key )
}

//获取hashKey的长度
func ( r *RedisMgr ) Hlen( hashKey string ) ( int, error )  {
	conn := r.GetConn()
	if conn.Err() != nil{
		return 0, conn.Err()
	}

	defer func() {
		conn.Close()
	}()

	return redis.Int( conn.Do("HLEN", hashKey ) )
}

//给hashKey里面指定的key建增加incrNum
//incrNum 必须为数字型
func ( r *RedisMgr) Hincrby ( hashKey string, key string, incrNum interface{} ) error  {
	switch incrNum.(type) {
		case int32, int, int64, int8, int16, float64, float32:
		default:
			return errors.New("参数incrNum必须为数字类型")
	}

	conn := r.GetConn()
	if conn.Err() != nil{
		return conn.Err()
	}

	defer func() {
		conn.Flush()
		conn.Close()
	}()


	return conn.Send("HINCRBY", hashKey, key, incrNum )

}

//给指定的key增加num
//num 必须为数字型
func ( r *RedisMgr) Incrnum( key string, num interface{} ) error {
	switch num.(type) {
		case int32, int, int64, int8, int16, float64, float32:
		default:
			return errors.New("参数num必须为数字类型")
	}

	conn := r.GetConn()
	if conn.Err() != nil{
		return conn.Err()
	}

	defer func() {
		conn.Flush()
		conn.Close()
	}()

	return  conn.Send("INCRBY", key, num )

}

// 设置有序集合
func ( r *RedisMgr ) Zset( key string, score interface{}, member string )  error {
	conn := r.GetConn()
	if conn.Err() != nil{
		return conn.Err()
	}

	defer func() {
		conn.Flush()
		conn.Close()
	}()
	return  conn.Send("ZADD", key, score, member )

}

//获取有序集合的数据
func ( r *RedisMgr ) Zrange( key string, start int, end int, desc string, withScores bool ) ( interface{}, error ) {
	conn := r.GetConn()
	if conn.Err() != nil{
		return nil, conn.Err()
	}

	defer func() {
		conn.Close()
	}()

	if strings.ToLower(desc) == "asc"{
		if withScores{
			return  conn.Do("ZRANGE", key, start, end, "WITHSCORES" )
		}else {
			return  conn.Do("ZRANGE", key, start, end )
		}
	}else{
		if withScores{
			return conn.Do("ZREVRANGE", key, start, end, "WITHSCORES" )
		}else{
			return conn.Do("ZREVRANGE", key, start, end )
		}
	}
}

//删除有序集合key里面的member成员
func ( r *RedisMgr) Zdel( key string, member string ) error {
	conn := r.GetConn()
	if conn.Err() != nil{
		return conn.Err()
	}

	defer func() {
		conn.Flush()
		conn.Close()
	}()

	return conn.Send( "ZREM", key, member )
}

//计算有序集合在指定分数范围内的长度
func ( r *RedisMgr ) Zcount( key string, minSorce int, maxSorce int ) (int, error )  {
	conn := r.GetConn()
	if conn.Err() != nil{
		return 0, conn.Err()
	}

	defer func() {
		conn.Close()
	}()

	return redis.Int( conn.Do( "ZCOUNT", key, minSorce, maxSorce ) )
}

//获取某个分数段的集合
func ( r *RedisMgr ) ZrangeByScore( key string, minSorce int, maxSorce int ) ( interface{}, error ) {
	conn := r.GetConn()
	if conn.Err() != nil{
		return nil, conn.Err()
	}

	defer func() {
		conn.Close()
	}()

	return conn.Do("ZRANGEBYSCORE", minSorce, maxSorce )
}

//获取成员member在有序集合key里面的排名
func ( r *RedisMgr ) Zrank( key string, member string, sort string )  ( int, error ) {
	conn := r.GetConn()
	if conn.Err() != nil{
		return 0, conn.Err()
	}

	defer func() {
		conn.Close()
	}()

	if strings.ToLower( sort ) == "asc"{
		return redis.Int( conn.Do("ZRANK", key, member ) )
	}else {
		return redis.Int( conn.Do("ZREVRANK", key, member ) )
	}
}

//关闭所有连接
func ( r *RedisMgr) Close()  {
	r.pool.Close()
}