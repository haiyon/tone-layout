package server

import (
	"sample/internal/conf"
	"sample/pkg/utils"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

// NewRegistrar - new consul registry.
func NewRegistrar(cr *conf.Registry) registry.Registrar {
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
