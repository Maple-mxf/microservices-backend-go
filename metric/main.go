package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/credentials"
)

var tracer trace.Tracer

func newExporter() (sdktrace.SpanExporter, error) {
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	return otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint("172.17.0.1:4317"),
		),
	)
	//if err != nil {
	//	panic(err)
	//}
	//return stdouttrace.New(
	//	stdouttrace.WithWriter(w),
	//	// Use human-readable output.
	//	stdouttrace.WithPrettyPrint(),
	//	// Do not print timestamps for the demo.
	//	stdouttrace.WithoutTimestamps(),
	//)
	//return exporter, nil
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
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)

}

func initTrace() {
	//ctx := context.Background()

	//f, err := os.OpenFile("traces.txt", os.O_RDWR, 0777)

	exp, err := newExporter()
	if err != nil {
		panic(err)
	}

	tp := newTraceProvider(exp)
	// Handle shutdown properly so nothing leaks.
	//defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	tracer = tp.Tracer("ExampleService")
}

func main() {
	initTrace()
	http.HandleFunc("/", httpHandler)

	log.Fatal(http.ListenAndServe(":8888", nil))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), "hello-span")
	defer span.End()

	w.Write([]byte("Hello"))

	time.Sleep(time.Second * 3)
}
