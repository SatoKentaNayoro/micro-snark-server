package data

import (
	nacos "github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"micro-snark-server/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewRegistrar, NewDiscovery)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	client, err := clients.NewNamingClient(vo.NacosClientParam{ClientConfig: conf.ClientConf, ServerConfigs: conf.Endpoints})
	if err != nil {
		panic(err)
	}
	return nacos.New(client)
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	client, err := clients.NewNamingClient(vo.NacosClientParam{ClientConfig: conf.ClientConf, ServerConfigs: conf.Endpoints})
	if err != nil {
		panic(err)
	}
	return nacos.New(client)
}
