package main

import (
	"flag"
	"os"
	"sample/internal/build"
	"sample/internal/conf"

	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"

	"sample/pkg/observes"
	"sample/pkg/utils"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Identifier is the identifier.
	Identifier = "sample.stocms.com"
	// Name is the name.
	Name = "sample"
	// flagconf is the config flag.
	flagconf string
	// hostname is the run host name
	hostname, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../config.yaml", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(Identifier),
		kratos.Name(Name),
		kratos.Version(build.Version),
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
	logger := log.With(log.NewStdLogger(os.Stdout),
		"identifier", Identifier,
		"name", Name,
		"version", build.Version,
		"built_at", build.BuiltAt,
		"hostname", hostname,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
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

	if err := observes.NewTracer(&observes.TracerOption{
		URL:         bc.Trace.Endpoint,
		Name:        Name,
		Version:     build.Version,
		Branch:      build.Branch,
		Revision:    build.Revision,
		Environment: bc.Server.Mode,
	}); utils.IsNotNil(err) {
		log.Fatalf("tracer.Init: %s", err)
	}

	if err := observes.NewSentry(&observes.SentryOptions{
		Dsn:         bc.Sentry.Endpoint,
		Name:        Name,
		Release:     build.Version,
		Environment: bc.Server.Mode,
	}); err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

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
