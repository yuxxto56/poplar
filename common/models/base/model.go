package base
import "github.com/astaxie/beego/orm"

type Model struct {
	TableName string
	o * orm.Ormer
}

func (m *Model) Table(table string) *Model{
     m.TableName = table
     return m
}

func (m *Model) Limit(start interface{},limit interface{}){

}


