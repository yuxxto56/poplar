package routers

import (
    "github.com/smallnest/rpcx/server"
    "log"
    "poplar/rpc/service"
)

func InitRpcRouters(rpcServer *server.Server) {
    if err := rpcServer.Register(new(service.Student),""); err != nil {
        log.Fatalf("failed to register rpcRouter: %v", err)
    }
}