package logics

import (
	"fmt"
	"poplar/common/models"
)

type StudentLogic struct{}

func (s *StudentLogic) GetAll() []map[string]interface{}{
	sdu := new(models.StudentModel).Init()
	result := sdu.GetAll()
	fmt.Println("lastSql:",sdu.Model.GetLastSql())
	return result
}