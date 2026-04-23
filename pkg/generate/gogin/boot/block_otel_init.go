//ff:func feature=gen-gogin type=generator control=selection topic=observability
//ff:what blockOtelInit — OpenTelemetry TracerProvider 초기화 (exporter/sampler/shutdown)

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockOtelInit emits OpenTelemetry tracer-provider bootstrap code into
// main.go when manifest.backend.observability.tracing.enabled is true.
// Inert (no imports, no lines) otherwise — the generator skips it via
// collectActiveBlocks and go.mod stays free of OTel deps.
//
// Three exporter flavors, each isolated so a project pulls in only the
// SDK bits it actually uses at runtime:
//
//   - "otlp":   go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc
//               (default gRPC → localhost:4317, Jaeger or otel-collector)
//   - "stdout": go.opentelemetry.io/otel/exporters/stdout/stdouttrace
//               (dev-mode — spans printed to stderr)
//   - "noop":   no exporter. The SDK + sampler still run, which exercises the
//               instrumentation path end-to-end without shipping any data.
//
// The TracerProvider is set globally (`otel.SetTracerProvider`) and the
// W3C TraceContext + Baggage propagators are registered so otelgin,
// otelsql, and ssac/pkg/queue all observe the same span tree. Shutdown
// is deferred inside main() with a 5s timeout bound to a fresh
// context.Background so ongoing export flushes are not killed by the
// bootstrap ctx cancellation.
//
// Env overrides (read inline so ops can flip tracing off or swap the
// collector endpoint without regeneration):
//
//	BACKEND_OBSERVABILITY_TRACING_ENABLED       (bool, default = manifest value)
//	BACKEND_OBSERVABILITY_TRACING_SERVICE_NAME  (string, default = manifest)
//	BACKEND_OBSERVABILITY_TRACING_EXPORTER      (otlp/stdout/noop)
//	BACKEND_OBSERVABILITY_TRACING_OTLP_ENDPOINT (host:port)
//	BACKEND_OBSERVABILITY_TRACING_SAMPLE_RATE   (float64 in [0,1])
//
// Ordering (collect_active_blocks): emitted BEFORE blockRouter so the
// gin middleware chain can reference the shared tracer; emitted AFTER
// blockServerStruct so server fields (if they later need a Tracer handle)
// can be assigned in-scope.
func blockOtelInit(fs *yongol.Fullstack, modulePath string) MainBlock {
	if !hasOtel(fs) {
		return MainBlock{Name: "otel-init"}
	}

	serviceName := otelServiceName(fs)
	exporter := otelExporter(fs)
	endpoint := otelOtlpEndpoint(fs)
	sampleRate := otelSampleRate(fs)

	imports := []string{
		`"go.opentelemetry.io/otel"`,
		`"go.opentelemetry.io/otel/propagation"`,
		`sdktrace "go.opentelemetry.io/otel/sdk/trace"`,
		`"go.opentelemetry.io/otel/sdk/resource"`,
		`semconv "go.opentelemetry.io/otel/semconv/v1.26.0"`,
	}

	lines := []string{
		fmt.Sprintf(`otelServiceName := envString("BACKEND_OBSERVABILITY_TRACING_SERVICE_NAME", %q)`, serviceName),
		fmt.Sprintf(`otelExporter := envString("BACKEND_OBSERVABILITY_TRACING_EXPORTER", %q)`, exporter),
		fmt.Sprintf(`otelSampleRate := envFloat("BACKEND_OBSERVABILITY_TRACING_SAMPLE_RATE", %v)`, sampleRate),
		`otelEnabled := envBool("BACKEND_OBSERVABILITY_TRACING_ENABLED", true)`,
		`var otelShutdown func(context.Context) error`,
		`if otelEnabled {`,
		`	var spanExporter sdktrace.SpanExporter`,
		`	switch otelExporter {`,
	}

	switch exporter {
	case "otlp":
		imports = append(imports, `otlpgrpc "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"`)
		lines = append(lines,
			`	case "otlp":`,
			fmt.Sprintf(`		otlpEndpoint := envString("BACKEND_OBSERVABILITY_TRACING_OTLP_ENDPOINT", %q)`, endpoint),
			`		exp, err := otlpgrpc.New(ctx,`,
			`			otlpgrpc.WithEndpoint(otlpEndpoint),`,
			`			otlpgrpc.WithInsecure(),`,
			`		)`,
			`		if err != nil { slog.Error("otel otlp exporter", "err", err); os.Exit(1) }`,
			`		spanExporter = exp`,
		)
	case "stdout":
		imports = append(imports, `"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"`)
		lines = append(lines,
			`	case "stdout":`,
			`		exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())`,
			`		if err != nil { slog.Error("otel stdout exporter", "err", err); os.Exit(1) }`,
			`		spanExporter = exp`,
		)
	case "noop":
		lines = append(lines,
			`	case "noop":`,
			`		spanExporter = nil`,
		)
	}

	lines = append(lines,
		`	default:`,
		`		slog.Error("otel exporter not supported by this build", "exporter", otelExporter)`,
		`		os.Exit(1)`,
		`	}`,
		`	res, err := resource.New(ctx, resource.WithAttributes(`,
		`		semconv.ServiceName(otelServiceName),`,
		`	))`,
		`	if err != nil { slog.Error("otel resource", "err", err); os.Exit(1) }`,
		`	tpOpts := []sdktrace.TracerProviderOption{`,
		`		sdktrace.WithResource(res),`,
		`		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(otelSampleRate)),`,
		`	}`,
		`	if spanExporter != nil {`,
		`		tpOpts = append(tpOpts, sdktrace.WithBatcher(spanExporter))`,
		`	}`,
		`	tp := sdktrace.NewTracerProvider(tpOpts...)`,
		`	otel.SetTracerProvider(tp)`,
		`	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(`,
		`		propagation.TraceContext{},`,
		`		propagation.Baggage{},`,
		`	))`,
		`	otelShutdown = tp.Shutdown`,
		`	slog.Info("otel tracing initialized", "service", otelServiceName, "exporter", otelExporter, "sample_rate", otelSampleRate)`,
		`}`,
		`defer func() {`,
		`	if otelShutdown == nil { return }`,
		`	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)`,
		`	defer cancel()`,
		`	if err := otelShutdown(shutdownCtx); err != nil {`,
		`		slog.Warn("otel shutdown", "err", err)`,
		`	}`,
		`}()`,
	)

	// Ensure `time` is available for the shutdown timeout even if no other
	// block already imported it.
	imports = append(imports, `"time"`)

	return MainBlock{
		Name:    "otel-init",
		Imports: imports,
		Lines:   lines,
	}
}
