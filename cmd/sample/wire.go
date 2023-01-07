//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"sample/internal/conf"
	"sample/internal/data"
	"sample/internal/server"
	"sample/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func wireApp(*conf.Server, *conf.Auth, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
