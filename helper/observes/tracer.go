package observes

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
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
	if opt == nil || opt.URL == "" {
		fmt.Println("Not exist tracer config...")
		return nil
	}

	// GRPC
	exp, err := otlptracegrpc.New(context.Background(), otlptracegrpc.WithEndpoint(opt.URL), otlptracegrpc.WithInsecure())
	if err != nil {
		return err
	}

	// HTTP
	// exp, err := otlptracehttp.New(context.Background(), otlptracehttp.WithEndpoint(opt.URL), otlptracehttp.WithInsecure())
	// if err != nil {
	// 	return err
	// }

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
