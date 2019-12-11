package models

import (
	"poplar/common/functions"
	"poplar/common/models/base"
)

//定义Model
type StudentModel struct {
	Model *base.Model
	Field StudentModelField
}
//定义Model字段
type StudentModelField struct {
	T_table string `default:"student"`
	F_id    string `default:"id"`
	F_name  string `default:"name"`
	F_age   string `default:"age"`
}

//新增
func (s *StudentModel) Insert(data map[string]interface{}) (int){
     result,_ := s.Model.Data(data).Insert()
	 return result
}

//获取单条数据
func (s *StudentModel) GetById(id int,field ...[]string) map[string]interface{}{
	if len(field) > 0{
		s.Model.Field(field[0])
	}
    rs := s.Model.Where(map[string]interface{}{
	"id":id,
   }).Find()
    return rs
}

func (s *StudentModel) GetAll() []map[string]interface{}{
    rs := s.Model.Limit(2,3).Select()
    return rs
}



func (s *StudentModel) Init() *StudentModel{
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table)
	return s
}










