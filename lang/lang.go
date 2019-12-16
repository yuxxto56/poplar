package lang

import (
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
)

//定义变量
var(
	loadLang *Load
)

//结构体
type Load struct {
	i18n *i18n.Locale
}

//获取lang
func (l *Load) GetLang(str string) string{
	return getLang(str)
}

//返回值
func getLang(str string) string{
	return loadLang.i18n.Tr(str)
}

//实例化
func init(){
	//加载语言包
	if err := i18n.SetMessage("zh-CN", "lang/lang_zh_cn.ini");err != nil{
		logs.Info("load lang/lang_zh_cn.ini Error")
	}
	loadLang = &Load{
		&i18n.Locale{Lang:"zh-CN"},
	}
}