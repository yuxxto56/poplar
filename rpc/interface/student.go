/**
 * @Author: Gosin
 * @Date: 2019/12/12 15:38
 */
package _interface

import "context"

type Student interface {
    GetAll(ctx context.Context, args *map[string]interface{}, reply *[]map[string]interface{}) error
}