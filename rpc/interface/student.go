/**
 * @Author: Gosin
 * @Date: 2019/12/12 15:38
 */
package _interface

import "context"
type Reply struct {
    A string
    B string
}

type Student interface {
    GetAll(ctx context.Context, args *Reply, reply *[]map[string]interface{}) error
}