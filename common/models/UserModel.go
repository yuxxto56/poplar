package models

import (
	"poplar/common/functions"
	"poplar/common/models/base"
)

//表名
const tables = "user"

//定义内部变量
var user *UserModel

//定义Model
type UserModel struct {
	Model *base.Model
	Field UserModelField
}
//定义Model字段
type UserModelField struct {
	F_id   string `field:"id"`
	F_name string `field:"name"`
	F_age  string `field:"age"`
}

func (s *UserModel) Insert(data map[string]interface{}) (int){
	result,_ := s.Model.Data(data).Insert()
	return result
}

//实例
func NewUserModel() *UserModel{
	return user
}

//初始化
func init(){
	user = &UserModel{
		Model:base.NewModel(tables),
	}
	functions.ReflectModel(&user.Field)
}











