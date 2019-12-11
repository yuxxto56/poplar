package toolLib

import (
	"context"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"net/url"
	"strings"
	"time"
)

var (
	RedisGlobMgr 		*RedisMgr = &RedisMgr{}
)


type RedisMgr struct {
	ctx         context.Context
	pool        *pools.ResourcePool

}

type RedisConn struct {
	redis.Conn
}


func (r *RedisConn) Close() {
	_ = r.Conn.Close()
}

func init()  {
	RedisInit( beego.AppConfig.String("redis.db"), RedisGlobMgr )
}

func RedisInit( db string, redism *RedisMgr )  {
	redism.ctx = context.Background()
	dns := fmt.Sprintf("redis://%s:%s/%s", beego.AppConfig.String("redis.host"), beego.AppConfig.String("redis.port"), db )
	capacity, err := beego.AppConfig.Int("redis.capacity")
	if err != nil{
		capacity = 50
	}
	maxCap, err := beego.AppConfig.Int("redis.maxCap")
	if err != nil{
		maxCap = 80
	}
	idleTimeout, err := beego.AppConfig.Int("redis.idleTimeout")
	if err != nil{
		idleTimeout = 600
	}
	redism.pool = RedisGlobMgr.newRedisPool(dns, capacity, maxCap, time.Duration(idleTimeout)*time.Second )
}

func (r *RedisMgr ) newRedisFactory(uri string) pools.Factory {
	return func() (pools.Resource, error) {
		return r.redisConnFromURI(uri)
	}
}

func (r *RedisMgr ) newRedisPool(uri string, capacity int, maxCapacity int, idleTimout time.Duration) *pools.ResourcePool {
	return pools.NewResourcePool(r.newRedisFactory(uri), capacity, maxCapacity, idleTimout)
}

func (r *RedisMgr ) redisConnFromURI(uriString string) (*RedisConn, error) {
	uri, err := url.Parse(uriString)
	if err != nil {
		return nil, err
	}

	var network string
	var host string
	var password string
	var db string
	var dialOptions []redis.DialOption

	switch uri.Scheme {
	case "redis", "rediss":
		network = "tcp"
		host = uri.Host
		if uri.User != nil {
			password, _ = uri.User.Password()
		}
		if len(uri.Path) > 1 {
			db = uri.Path[1:]
		}
	case "unix":
		network = "unix"
		host = uri.Path
	default:
		return nil, errors.New("非法的Redis链接")
	}
	conn, err := redis.Dial(network, host, dialOptions...)
	if err != nil {
		return nil, err
	}

	if password != "" {
		_, err := conn.Do("AUTH", password)
		if err != nil {
			conn.Close()
			return nil, err
		}
	}

	if db != "" {
		_, err := conn.Do("SELECT", db)
		if err != nil {
			conn.Close()
			return nil, err
		}
	}

	return &RedisConn{Conn: conn}, nil
}


func (r *RedisMgr) GetConn() (*RedisConn, error) {
	resource, err := r.pool.Get(r.ctx)

	if err != nil {
		return nil, err
	}
	return resource.(*RedisConn), nil
}

func (r *RedisMgr ) PutConn(conn *RedisConn) {
	r.pool.Put(conn)
}


//向一个key[队列]的尾部添加一个元素
func (r *RedisMgr) Rpush( key string, data interface{} ) error {
	conn, err := r.GetConn()
	if err != nil{
		return err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()
	return  conn.Send("RPUSH", key, data )
}

//向一个key[队列]的头部添加一个元素
func (r *RedisMgr) Lpush( key string, data interface{} ) error {
	conn, err := r.GetConn()
	if err != nil{
		return err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()
	return  conn.Send("LPUSH", key, data )
}

//取出队列中第一个key取元素值
func (r *RedisMgr) Lpop( key string ) ( interface{}, error )   {
	conn, err := r.GetConn()
	if err != nil{
		return nil, err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()
	return  conn.Do("LPOP", key )
}

//返回名称为key的list中start至end之间的元素（end为 -1 ，返回所有）
func ( r *RedisMgr ) Lrange(key string, start int, end int )  ( interface{}, error )  {
	conn, err := r.GetConn()
	if err != nil{
		return nil, err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()
	return  conn.Do("LRANGE", key, start, end )
}

//获取队列长度
func ( r *RedisMgr ) Llen( key string ) (int , error) {
	conn, err := r.GetConn()
	if err != nil{
		return 0,err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()
	len, err := redis.Int( conn.Do("LLEN", key) )
	if err != nil{
		return 0,err
	}
	return len, nil
}

// 判断一个key集合里是否存在某个value值，存在返回True
func ( r *RedisMgr ) Scontains(key string, data interface{}) (bool, error)  {
	conn, err := r.GetConn()
	if err != nil{
		return false,err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()

	return  redis.Bool( conn.Do("SISMEMBER", key, data ) )
}

//向集合添加元素
func ( r *RedisMgr) Sadd(key string, data interface{})  error {
	conn, err := r.GetConn()
	if err != nil{
		return err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()

	return conn.Send("SADD", key, data )
}


//返回key集合所有的元素
func ( r *RedisMgr )  Smembers(key string) (interface{}, error){
	conn, err := r.GetConn()
	if err != nil{
		return nil,err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()
	return  conn.Do("SMEMBERS", key )
}

//在key集合中移除指定的元素
func ( r *RedisMgr ) Srem( key string, data interface{} ) error {
	conn, err := r.GetConn()
	if err != nil{
		return err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()

	return  conn.Send("SREM", key, data)
}

//删除指定的key
func ( r *RedisMgr ) Clear( key string) error  {
	conn, err := r.GetConn()
	if err != nil{
		return err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()

	return conn.Send("DEL", key )
}

//设置数据
func ( r *RedisMgr) Set(key string, data interface{} ) error  {
	conn, err := r.GetConn()
	if err != nil{
		return err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()

	return  conn.Send("SET", key, data )
}

//获取数据
func ( r *RedisMgr) Get(key string ) (interface{}, error)  {
	conn, err := r.GetConn()
	if err != nil{
		return "",err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()


	return  conn.Do("GET", key )

}

//设置某个hashKey名称的下的keyvalue值
func ( r *RedisMgr ) Hset( hashKey string, key string, data interface{} ) error {
	conn, err := r.GetConn()
	if err != nil{
		return err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()

	return conn.Send("HSET", hashKey, key, data )

}

//得到某个hashKey名称下的key信息
func (r *RedisMgr ) Hget( hashKey string, key string ) ( interface{}, error )  {
	conn, err := r.GetConn()
	if err != nil{
		return nil,err
	}

	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()

	return conn.Do("HGET", hashKey, key )
}

//删除haskKey下面的key建
func ( r *RedisMgr) Hdel( hashKey string, key string ) error  {
	conn, err := r.GetConn()
	if err != nil{
		return err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()

	return  conn.Send("HDEL", hashKey, key )
}

//获取hashKey的长度
func ( r *RedisMgr ) Hlen( hashKey string ) ( int, error )  {
	conn, err := r.GetConn()
	if err != nil{
		return 0, err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
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

	conn, err := r.GetConn()
	if err != nil{
		return err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
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

	conn, err := r.GetConn()
	if err != nil{
		return err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()

	return  conn.Send("INCRBY", key, num )

}

// 设置有序集合
func ( r *RedisMgr ) Zset( key string, score interface{}, member string )  error {
	conn, err := r.GetConn()
	if err != nil{
		return err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()
	return  conn.Send("ZADD", key, score, member )

}

//获取有序集合的数据
func ( r *RedisMgr ) Zrange( key string, start int, end int, desc string, withScores bool ) ( interface{}, error ) {
	conn, err := r.GetConn()
	if err != nil{
		return nil, err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
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
	conn, err := r.GetConn()
	if err != nil{
		return err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()

	return conn.Send( "ZREM", key, member )
}

//计算有序集合在指定分数范围内的长度
func ( r *RedisMgr ) Zcount( key string, minSorce int, maxSorce int ) (int, error )  {
	conn, err := r.GetConn()
	if err != nil{
		return 0, err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()

	return redis.Int( conn.Do( "ZCOUNT", key, minSorce, maxSorce ) )
}

//获取某个分数段的集合
func ( r *RedisMgr ) ZrangeByScore( key string, minSorce int, maxSorce int ) ( interface{}, error ) {
	conn, err := r.GetConn()
	if err != nil{
		return nil, err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()

	return conn.Do("ZRANGEBYSCORE", minSorce, maxSorce )
}

//获取成员member在有序集合key里面的排名
func ( r *RedisMgr ) Zrank( key string, member string, sort string )  ( int, error ) {
	conn, err := r.GetConn()
	if err != nil{
		return 0, err
	}
	defer func() {
		conn.Flush()
		r.PutConn(conn)
	}()

	if strings.ToLower( sort ) == "asc"{
		return redis.Int( conn.Do("ZRANK", key, member ) )
	}else {
		return redis.Int( conn.Do("ZREVRANK", key, member ) )
	}
}