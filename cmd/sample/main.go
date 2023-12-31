package main

import (
	"flag"
	"fmt"
	"os"
	"sample/helper/config"
	"sample/internal"
	"sample/internal/conf"

	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"

	"sample/helper/observes"
	"sample/helper/utils"

	"github.com/go-kratos/kratos/v2"
	kc "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name.
	Name = "sample"
	// flagConf is the config flag.
	flagConf string
	// hostname is the run host name
	hostname, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagConf, "conf", "config.yaml", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(internal.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(rr),
	)
}

func main() {
	flag.Parse()
	c := kc.New(
		kc.WithSource(
			config.Loader(flagConf),
		),
	)
	defer c.Close()

	if err := c.Load(); utils.IsNotNil(err) {
		log.Fatalf("config.load: %s", err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); utils.IsNotNil(err) {
		log.Fatalf("config.bootstrap: %s", err)
	}

	var rc conf.Registry
	if err := c.Scan(&rc); utils.IsNotNil(err) {
		log.Fatalf("config.registry: %s", err)
	}

	// application run name
	if bc.Server.Mode != "" && bc.Server.Name != "" {
		Name = fmt.Sprintf("%s.%s", bc.Server.Mode, bc.Server.Name)
	}
	// init tracer
	if err := observes.NewTracer(&observes.TracerOption{
		URL:         bc.Trace.Endpoint,
		Name:        Name,
		Version:     internal.Version,
		Branch:      internal.Branch,
		Revision:    internal.Revision,
		Environment: bc.Server.Mode,
	}); utils.IsNotNil(err) {
		log.Fatalf("tracer.Init: %s", err)
	}
	// init sentry
	if err := observes.NewSentry(&observes.SentryOptions{
		Dsn:         bc.Sentry.Endpoint,
		Name:        Name,
		Release:     internal.Version,
		Environment: bc.Server.Mode,
	}); err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// init logger
	logger := log.With(log.NewStdLogger(os.Stdout),
		"name", Name,
		"version", internal.Version,
		"built_at", internal.BuiltAt,
		"hostname", hostname,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)

	app, cleanup, err := wireApp(bc.Server, bc.Auth, &rc, bc.Data, logger)
	if utils.IsNotNil(err) {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); utils.IsNotNil(err) {
		panic(err)
	}
}
