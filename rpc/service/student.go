/**
 * @Author: Gosin
 * @Date: 2019/12/12 15:40
 */
package service

import (
    "context"
)

type Student struct {
}
func (s *Student)GetAll(ctx context.Context, args *[]map[string]interface{}, reply *[]map[string]interface{}) error {
    //(*reply) = new(logics.StudentLogic).GetAll()
    return nil
}