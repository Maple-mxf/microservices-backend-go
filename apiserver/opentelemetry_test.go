package main

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"io"
	"os"
	"testing"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

var tracer trace.Tracer

func newExporter(w io.Writer) (sdktrace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

// newTraceProvider
func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	r, err := resource.Merge(resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("SampleService")))

	if err != nil {
		panic(err)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r))

}

func TestTelemetryAPI(t *testing.T) {
	ctx := context.Background()

	f, err := os.Create("traces.txt")
	defer f.Close()

	if err != nil {
		panic(err)
	}
	exp, err := newExporter(f)
	if err != nil {
		panic(err)
	}

	tp := newTraceProvider(exp)
	// Handle shutdown properly so nothing leaks.
	defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	tracer = tp.Tracer("ExampleService")

}

//
//func httpHandler(w http.ResponseWriter, r *http.Request) {
//	ctx, span := tracer.Start(r.Context(), "hello-span")
//	defer span.End()
//
//	// do some work to track with hello-span
//}
