package base

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strconv"
	"strings"
)
//定义变量
var (
	T_PREFIX = beego.AppConfig.String("db.prefix")
)

//定义Model结构体
type Model struct {
	table string
	o orm.Ormer
	limit []interface{}
	orderBy []string
	where map[string]interface{}
	data  map[string]interface{}
	field  string
	sql    string
}

//操作另一张表,表名不需要扩展
func (m *Model) Table(table string) *Model{
     m.table = T_PREFIX+table
     return m
}

//条件查询
//map[string]interface{"id":1,"name":[]string{"in","1,2,3"}})
func (m *Model) Where(param map[string]interface{}) *Model{
    if len(param) == 0{
    	return m
	}
	m.where = make(map[string]interface{})
	for k,v := range param{
		if reflect.TypeOf(v).String() == "[]string"{
			vs := v.([]string)
			if strings.Contains(vs[1],","){
				vs[1] = "("+vs[1]+")"
			}
			m.where[k+"_"+strings.ToLower(vs[0])] = vs[1]
		}else{
			m.where[k] = v
		}
    }
	return m
}

//设置查询范围
//使用示例 Limit(10) limit(0,10)
func (m *Model) Limit(start interface{},limit ...interface{}) *Model{
	m.limit = make([]interface{},0)
    if len(limit) == 0{
    	m.limit = append(m.limit,start)
	}else{
		m.limit = append(m.limit,start)
		m.limit = append(m.limit,limit[0])
	}
	return m
}
//设置排序
//使用示例 OrderBy("id asc","age desc")
func (m *Model) OrderBy(params ...string) *Model{
	 if len(params) == 0{
	 	return m
	 }
	 m.orderBy = make([]string,len(params))
	 for k,v := range params{
	     v = strings.ToLower(v)
		 m.orderBy[k] = v
	 }
	 return m
}

//查询的字段
//使用示例 Field([]string{"name","age"})
func (m *Model) Field(param ...[]string) *Model{
	if len(param) == 0 {
		m.field = "*"
	}else{
		m.field = strings.TrimRight(strings.Join(param[0],","),",")
	}
	return m
}

//存储新增、更新数据
//使用示例
//maps := make(map[string]interface{})
//maps["name"] = "lidazhao"
//maps["age"]  = 21
//Data(maps)
func (m *Model) Data(param map[string]interface{}) *Model{
	m.data = make(map[string]interface{})
	m.data = param
	return m
}


//新增数据
func (m *Model) Insert()(int,error){
	//分析参数
	if len(m.data) == 0{
		return 0,nil
	}
	var colsName,colsValue = "",""
	for i,v := range m.data{
        colsName += "`"+i+"`"+","
		//如果为整型则转字符串类型
		if vs, p := v.(int); p {
			v = strconv.Itoa(vs)
		}
        colsValue += "'"+v.(string)+"'"+","
	}
	colsName  = strings.TrimRight(colsName,",")
	colsValue = strings.TrimRight(colsValue,",")
	// 组合数据写入SQL
	sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s);",m.table,colsName,colsValue)
	m.sql = sql
	retData, err := m.o.Raw(sql).Exec()
	if err != nil {
		return 0, nil
	}
	LastId, err := retData.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(LastId), err
}

//查询多条数据
func (m *Model) Select() []map[string]interface{}{
	var field string
	if m.field == ""{
		field = "*"
	}else{
		field = m.field
	}
    where := m.whereString()
	var orderBy string
	if len(m.orderBy)>0{
		for _,v := range m.orderBy{
			orderBy += v+","
		}
		orderBy = " ORDER BY "+strings.TrimRight(orderBy,",")
	}
	var limit string
	if len(m.limit)>0{
		if len(m.limit) == 1{
			    limit = strconv.Itoa(m.limit[0].(int))
		}else{
			    limit = strconv.Itoa(m.limit[0].(int)) + "," + strconv.Itoa(m.limit[1].(int))
		}
		limit = " LIMIT "+limit
	}
	sql := fmt.Sprintf("SELECT %s FROM %s%s%s%s",field,m.table,where,orderBy,limit)
	m.sql = sql
	var res []orm.Params
	_,err := m.o.Raw(sql).Values(&res)
	if err != nil{
		logs.Error("Sql:",sql," Error,",err.Error())
	}
	var maps = make([]map[string]interface{},len(res))
	if len(res)>0{
		for i,v := range res{
			maps[i] = v
		}
	}
	return maps
}

//查询单条数据
func (m *Model) Find() map[string]interface{}{
   var field string
   if m.field == ""{
	   field = "*"
   }else{
   	   field = m.field
   }
   where := m.whereString()
   sql := fmt.Sprintf("SELECT %s FROM %s%s LIMIT 1",field,m.table,where)
   m.sql = sql
   var res []orm.Params
   _,err := m.o.Raw(sql).Values(&res)
   if err != nil{
   	  logs.Error("Sql:",sql," Error,",err.Error())
   }
   if(len(res) == 0){
   	  return make(map[string]interface{})
   }
   return res[0]
}

//更新
func (m *Model) Update() (int,error){
	//分析参数
	if len(m.data) == 0{
		return 0,nil
	}
	var updateStr string
	for i,v := range m.data{
		//如果为整型则转字符串类型
		if vs, p := v.(int); p {
			v = strconv.Itoa(vs)
		}
		updateStr += i+"="+"'"+v.(string)+"'"+","
	}
	updateStr = strings.TrimRight(updateStr,",")
	where := m.whereString()
	sql := fmt.Sprintf("UPDATE %s SET %s%s",m.table,updateStr,where)
	m.sql = sql
	sqlSource,err := m.o.Raw(sql).Exec()
	if err != nil{
		logs.Error("sql:",sql,"Error ",err.Error())
		return 0,nil
	}
	num,_ := sqlSource.RowsAffected()
	return int(num),err
}

//物理删除
func (m *Model) Delete()(int,error){
	where := m.whereString()
	if where == ""{
		return 0,nil
	}
	sql := fmt.Sprintf("DELETE FROM %s%s",m.table,where)
	m.sql = sql
	sqlSource,err := m.o.Raw(sql).Exec()
	if err != nil{
		logs.Error("sql:",sql,"Error ",err.Error())
		return 0,nil
	}
	num,_ := sqlSource.RowsAffected()
	return int(num),err
}

//统计数据
//使用示例
//Count() 或 Count("id") //id为字段名
func (m *Model) Count(param ...string) (int){
	co := "*"
	if len(param) != 0 {
		co = param[0]
	}
	where := m.whereString()
	sql := fmt.Sprintf("SELECT COUNT(%s) FROM %s%s",co,m.table,where)
	m.sql = sql
	var maps []orm.Params
	_,err := m.o.Raw(sql).Values(&maps)
	if err != nil{
		logs.Error("Sql:",sql," Error,",err.Error())
	}
	num,_:=strconv.Atoi(maps[0]["COUNT("+co+")"].(string))
	return num
}

//事务开始
func (m *Model) Begin()(*Model){
	err := m.o.Begin()
	if err != nil{
		logs.Error("Begin Error",err.Error())
	}
	return m
}

//事务提交
func (m *Model) Commit()(*Model){
	err := m.o.Commit()
	if err != nil{
		logs.Error("Commit Error",err.Error())
	}
	return m
}

//事务回滚
func (m *Model) RollBack()(*Model){
	err := m.o.Rollback()
	if err != nil{
		logs.Error("RollBack Error",err.Error())
	}
	return m
}

//打印sql语句
//使用示例
//GetLastSql() 返回sql语句 GetLastSql(true) 打印控制台
func (m *Model) GetLastSql(param ...bool) (string){
	 var isPrint bool
	 if len(param) != 0{
	 	 isPrint = true
	 }
	 if isPrint{
	 	fmt.Println("sql:",m.sql)
	 	return ""
	 }
	 return m.sql
}


//组织where字符串
func (m *Model) whereString()(string){
	var where string = ""
	if len(m.where) != 0{
		for i,v := range m.where{
			//如果为整型则转字符串类型
			if vs, p := v.(int); p {
				v = strconv.Itoa(vs)
			}
			if strings.Contains(i,"_"){
				is := strings.Split(i,"_")
				where += is[0]+" "+is[1]+" "+v.(string)+" AND "
			}else{
				where += i+"="+"'"+v.(string)+"'"+" AND "
			}
		}
		where = " WHERE "+strings.TrimRight(where," AND")
	}
	return where
}

//实例化Model引用
//@param string table 表名称
//使用示例
//NewModel("student")
func NewModel(table string) *Model{
	ormers := orm.NewOrm()
	return &Model{
		table:T_PREFIX+table,
		o:ormers,
	}
}
