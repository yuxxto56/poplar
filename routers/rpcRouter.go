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

// func addRegistryPlugin(s *server.Server) {
// r := &serverplugin.EtcdRegisterPlugin{
//     ServiceAddress: "tcp@0.0.0.0:8088",
//     EtcdServers:    []string{"0.0.0.0:8888","0.0.0.0:8888"},
//     BasePath:       beego.AppConfig.String("appname"),
//     Metrics:        metrics.NewRegistry(),
//     UpdateInterval: time.Minute,
// }
// err := r.Start()
// if err != nil {
//     log.Fatal(err)
// }
// s.Plugins.Add(r)
// }