/**
 * @Author: Gosin
 * @Date: 2019/12/12 15:38
 */
package poplar

import "context"
type Args struct {
    A string
    B string
}

type Student interface {
    GetAll1(ctx context.Context, args *Args, reply *[]map[string]interface{}) error
    GetAll2(ctx context.Context, args *map[string]interface{}, reply *[]map[string]interface{}) error
}