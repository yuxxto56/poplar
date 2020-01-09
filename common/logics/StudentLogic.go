package logics

import (
	"errors"
	"poplar/common/models"
	_const "poplar/lang/const"
)

type StudentLogic struct{

}

func (s *StudentLogic) GetAll() ([]map[string]interface{},error){
	sdu := new(models.StudentModel).Init()
	result := sdu.GetAll()
	if len(result) == 0{
		err :=errors.New(_const.Hi)
		return nil,err
	}
	return result,nil
}