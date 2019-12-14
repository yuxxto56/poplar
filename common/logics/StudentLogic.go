package logics

import (
	"errors"
	"poplar/common/models"
	"poplar/lang"
	_const "poplar/lang/const"
)

type StudentLogic struct{
	lang.Load
}

func (s *StudentLogic) GetAll() ([]map[string]interface{},error){
	sdu := new(models.StudentModel).Init()
	result := sdu.GetAll()
	if len(result)>0{
		err :=errors.New(s.GetLang(_const.Hi))
		return nil,err
	}
	return result,nil
}