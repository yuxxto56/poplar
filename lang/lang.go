package lang

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
)

var(
	load *Load
)

type Load struct {
	i18n *i18n.Locale
}

func (l *Load) GetLang(str string) string{
	fmt.Println(l.i18n.Lang)
	return l.i18n.Tr(str)
}

//实例化
func init(){
	//加载语言包
	if err := i18n.SetMessage("zh-CN", "lang/locale_zh-CN.ini");err != nil{
		logs.Info("load lang/lang_zh.ini Error")
	}
	load = &Load{
		&i18n.Locale{Lang:"zh-CN"},
	}
}