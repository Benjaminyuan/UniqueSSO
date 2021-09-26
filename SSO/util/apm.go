package util

import (
	"context"
	"unique/jedi/conf"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	Tracer trace.Tracer
)

func SetupTracing() (func(ctx context.Context) error, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.SSOConf.APM.ReporterBackground)))
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		// Always be sure to batch in production.
		sdktrace.WithBatcher(exp),
		// Record information about this application in an Resource.
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(conf.SSOConf.Application.Name),
			attribute.String("environment", conf.SSOConf.Application.Mode),
		)),
	)

	otel.SetTracerProvider(tp)
	Tracer = otel.GetTracerProvider().Tracer(conf.SSOConf.Application.Name)
	return func(ctx context.Context) error {
		return tp.Shutdown(ctx)
	}, nil
}
