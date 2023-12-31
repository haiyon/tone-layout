package config

import (
	"net/url"

	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hashicorp/consul/api"
)

// Loader is handle config loader
func Loader(c string) config.Source {
	source := file.NewSource(c)

	if u, err := url.Parse(c); err != nil || u.Host == "" {
		return source
	}

	u, err := url.Parse(c)
	if err != nil {
		log.Fatalf("failed to parse url: %s", err)
	}

	cc, err := api.NewClient(&api.Config{
		Address: u.Host,
	})
	if err != nil {
		log.Fatalf("failed to create new client: %s", err)
		panic(err)
	}

	cs, err := consul.New(cc, consul.WithPath(u.Path))
	// The file suffix needs to be marked, needs to adapt the file suffix to read the configuration.
	if err != nil {
		log.Fatalf("failed to create new consul: %s", err)
		panic(err)
	}

	return cs
}
