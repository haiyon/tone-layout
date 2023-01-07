package observes

import (
	"fmt"

	"sample/pkg/utils"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type TracerOption struct {
	URL         string
	Name        string
	Version     string
	Branch      string
	Revision    string
	Environment string
}

// NewTracer is the register tracer
func NewTracer(opt *TracerOption) error {
	// if not exist tracer config, break
	if utils.IsNil(opt) {
		fmt.Println("Not exist tracer config...")
		return nil
	}
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(opt.URL)))
	if err != nil {
		return err
	}

	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(1.0))),
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(opt.Name),
			attribute.String("version", opt.Version),
			attribute.String("branch", opt.Branch),
			attribute.String("revision", opt.Revision),
			attribute.String("environment", opt.Environment),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return nil
}
