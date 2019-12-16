package poplar

import (
    "context"
    "poplar/rpc/interface/poplar"
)

type Student struct {
    Poplarbase
}
func (student *Student)Init() *Student {
    student.SetAddress()
    student.ServicePath = "Student"
    return student
}
func (student *Student)GetAll(ctx context.Context, args *poplar.Args, reply *[]map[string]interface{}) error {
    return student.GetXClient().Call(ctx, "GetAll1", args, reply)
}
func (student *Student)GetUserAll(args *map[string]interface{}) (*[]map[string]interface{},error) {
    reply := &[]map[string]interface{}{}
    return reply,student.GetXClient().Call(context.Background(), "GetAll2", args, reply)
}