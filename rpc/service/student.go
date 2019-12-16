/**
 * @Author: Gosin
 * @Date: 2019/12/12 15:40
 */
package service

import (
    "context"
    "poplar/common/logics"
    "poplar/rpc/interface/poplar"
)

type Student struct {
}
func (s *Student)GetAll1(ctx context.Context, args *poplar.Args, reply *[]map[string]interface{}) (err error) {
    (*reply),err = new(logics.StudentLogic).GetAll()
    return err
}
func (s *Student)GetAll2(ctx context.Context, args *map[string]interface{}, reply *[]map[string]interface{}) (err error) {
    (*reply),err = new(logics.StudentLogic).GetAll()
    return err
}