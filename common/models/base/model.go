package base

import "github.com/astaxie/beego/orm"

type Model struct {
	TableName string
	o orm.Ormer
}

func (m *Model) Table(table string) *Model{
     m.TableName = table
     return m
}

func (m *Model) Limit(start interface{},limit ...interface{}){
    if len(limit) == 0{
       // m.o.QueryTable()
       //m.o.q
	}
}

func (m *Model) Init(table string) *Model{
   return &Model{
	   TableName: table,
	   o:         orm.NewOrm(),
   }
}


