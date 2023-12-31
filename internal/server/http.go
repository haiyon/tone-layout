package server

import (
	iV1 "sample/api/interface/v1"
	"sample/helper/utils"
	"sample/internal/conf"
	"sample/internal/service"

	"github.com/gorilla/handlers"

	"github.com/go-kratos/sentry"

	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/validate"

	"github.com/go-kratos/swagger-api/openapiv2"

	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/selector"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer - new HTTP server.
func NewHTTPServer(c *conf.Server, svc *service.Service, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			sentry.Server(),
			logging.Server(logger),
			metadata.Server(),
			selector.Server().Match(matchWhiteList).Build(),
			getOperation,
			validate.Validator(),
		),
	}
	if utils.IsNotEmpty(c.Http.Network) {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if utils.IsNotEmpty(c.Http.Addr) {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	opts = append(opts,
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)))
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/q", openapiv2.NewHandler())
	iV1.RegisterPostHTTPServer(srv, svc)
	return srv
}
