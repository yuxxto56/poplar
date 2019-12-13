package poplar

import (
    "github.com/astaxie/beego"
    "poplar/rpc/client"
)

type Poplarbase struct {
    client.Baseclient
}
func (p *Poplarbase)SetAddress() {
    p.Address = beego.AppConfig.String("rpc.poplar.address")
}