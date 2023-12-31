package server

import (
	"sample/helper/utils"
	"sample/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

// NewDiscovery - new consul discovery.
func NewDiscovery(cr *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = cr.Consul.Address
	c.Scheme = cr.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if utils.IsNotNil(err) {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
