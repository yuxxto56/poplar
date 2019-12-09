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

func (s *StudentModel) Insert(data map[string]interface{}) (int){
     result,_ := s.Model.Data(data).Insert()
	 return result
}

//实例
func NewStudentModel() *StudentModel{
	models := &StudentModel{}
	functions.ReflectModel(&models.Field)
	models.Model = base.NewModel(models.Field.T_table)
	return models
}











