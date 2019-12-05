package base

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
)
//定义常量
var T_PREFIX = beego.AppConfig.String("db::prefix")

//定义Model结构体
type Model struct {
	table string
	o orm.Ormer
	limit []interface{}
	orderBy []string
	where map[string]interface{}
}

//操作另一张表,表名不需要扩展
func (m *Model) Table(table string) *Model{
     m.table = T_PREFIX+table
     return m
}

//map[string]interface{"id":1,"name":[]string["in","1,2,3"]})
func (m *Model) Where(param map[string]interface{}) *Model{
    if len(param) == 0{
    	return m
	}
	m.where = make(map[string]interface{})
	for k,v := range param{
		if reflect.TypeOf(v).String() == "[]string"{
			vs := v.([]string)
			m.where[k+"__"+strings.ToLower(vs[0])] = strings.Split(vs[1],",")
		}else{
			m.where[k] = v
		}
    }
	return m
}

//设置查询范围
//使用示例 Limit(10) limit(0,10)
func (m *Model) Limit(start interface{},limit ...interface{}) *Model{
    if len(limit) == 0{
	    m.limit[0] = start
	}else{
		m.limit[0] = start
		m.limit[1] = limit[0]
	}
	return m
}
//设置排序
//使用示例 OrderBy("id asc","age desc")
func (m *Model) OrderBy(params ...string) *Model{
	 if len(params) == 0{
	 	return m
	 }
	 for k,v := range params{
	     v = strings.ToLower(v)
	 	 if strings.Contains(v,"asc"){
              m.orderBy[k] = strings.TrimSpace(strings.Replace(v,"asc","",1))
		 }
		 if strings.Contains(v,"desc"){
			 m.orderBy[k] = "-"+strings.TrimSpace(strings.Replace(v,"desc","",1))
		 }
	 }
	 return m
}

//新增数据
//使用示例
func (m *Model) Insert(param map[string]interface{})(int,error){
	// 自定待创建的函数和参数
	insertCols, insertArgs := "", ""
	for k, v := range param {
		// 数据列只能为string类型
		if insertArgs == "" {
			insertArgs += fmt.Sprintf("%s", k)
		} else {
			insertArgs += fmt.Sprintf(",%s", k)
		}
		// 判断数据类型,类型断言判断
		switch v.(type) {
		case int:
			if insertArgs == "" {
				insertArgs += fmt.Sprintf("%d", v)
			} else {
				insertArgs += fmt.Sprintf(",%d", v)
			}
		case string:
			if insertArgs == "" {
				insertArgs += fmt.Sprintf("'%s'", v)
			} else {
				insertArgs += fmt.Sprintf(",'%s'", v)
			}
		case float64:
			if insertArgs == "" {
				insertArgs += fmt.Sprintf("%f", v)
			} else {
				insertArgs += fmt.Sprintf(",%f", v)
			}
		default:
			if insertArgs == "" {
				insertArgs += fmt.Sprintf("%v", v)
			} else {
				insertArgs += fmt.Sprintf(",%v", v)
			}
		}
	}
	// 组合数据写入SQL
	insertSql := fmt.Sprintf("INSERT INTO %v(%v) VALUES (%v);",m.table,insertCols, insertArgs)
	retData, err := m.o.Raw(insertSql).Exec()
	if err != nil {
		return 0, nil
	}
	LastId, err := retData.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(LastId), err
}

func (m *Model) Find(){

}


func NewModel(table string) *Model{
    ormer := orm.NewOrm()
    return &Model{
		table: T_PREFIX+table,
		o:         ormer,
	}
}

func init(){
	setDriver()
}

