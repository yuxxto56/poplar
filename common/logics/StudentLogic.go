package logics

import (
	"fmt"
	"poplar/common/models"
	"poplar/lang"
)

type StudentLogic struct{
	lang.Load
}

func (s *StudentLogic) GetAll() []map[string]interface{}{
	sdu := new(models.StudentModel).Init()
	result := sdu.GetAll()
	fmt.Println("lang:",s.GetLang("hi"))
	//fmt.Println("lastSql:",sdu.Model.GetLastSql())
	return result
}