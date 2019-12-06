package models

import (
	"poplar/common/models/base"
)

//表名
const table = "student"

//定义内部变量
var student *StudentModel

//定义Model
type StudentModel struct {
	Model *base.Model
	Field StudentModelField
}
//定义Model字段
type StudentModelField struct {
	F_id   string
	F_name string
	F_age  string
}

//实例
func NewStudentModel() *StudentModel{
	return student
}

//初始化
func init(){
	student = &StudentModel{
		Model:base.NewModel(table),
		Field:StudentModelField{
			F_id:"id",
			F_name:"name",
			F_age:"age",
		},
	}
}











