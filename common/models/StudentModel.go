package models

import "poplar/common/models/base"

//表名
const Table = "student"

//数据字段
const (
	F_id   = "id"
	F_name = "name"
	F_age  = "age"
)

type StudentModel struct {
	//继承Model类
	base.Model
}


//func (s *StudentModel) Insert(){
	//s.Table(Table).Data().
//}










