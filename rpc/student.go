/**
 * @Author: Gosin
 * @Date: 2019/12/12 15:40
 */
package rpc

import (
    "context"
    "poplar/common/logics"
)

type Student struct {
}
func (s *Student)GetAll(ctx context.Context, args *Reply, reply *[]map[string]interface{}) error {
    (*reply) = new(logics.StudentLogic).GetAll()
    return nil
}