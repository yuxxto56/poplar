package client

import (
    "github.com/smallnest/rpcx/client"
)

type Baseclient struct {
    Address string
    ServicePath string
    discovery client.ServiceDiscovery
    xClient client.XClient
}
func (cli *Baseclient)GetDiscovery() (client.ServiceDiscovery) {
    if cli.discovery == nil {
        cli.discovery = client.NewPeer2PeerDiscovery(cli.Address, "")
    }
    return cli.discovery
}

func (cli *Baseclient)GetXClient() (client.XClient) {
    if cli.xClient == nil {
        cli.xClient = client.NewXClient(cli.ServicePath, client.Failtry, client.RandomSelect, cli.GetDiscovery(), client.DefaultOption)
    }
    return cli.xClient
}

func (cli *Baseclient)Close() {
    cli.GetXClient().Close()
}