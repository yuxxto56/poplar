package toolLib

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/yunge/sphinx"
)

type SphinxMgr struct {
	Client *sphinx.Client
}

//建立sphinx客户端链接
func ( s *SphinxMgr) GetInstance()  ( *SphinxMgr, error ) {
	host := beego.AppConfig.String("sphinx.host")
	if host == ""{
		return s, errors.New("host 错误")
	}
	port, err := beego.AppConfig.Int("sphinx.port")
	if err != nil{
		port = 9312
	}
	s.Client = sphinx.NewClient().SetServer( host, port ).SetConnectTimeout(5000 )
	if err := s.Client.Error(); err != nil {
		return  s, err
	}
	return s, nil
}

//销毁
func ( s *SphinxMgr ) Destruct()  {
	s.Client.Close()
	s = nil
}
