//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what otelTailLines — exporter 분기 이후 공통으로 따라붙는 TracerProvider 초기화 + shutdown

package boot

// otelTailLines returns the shared TracerProvider + propagator + shutdown
// wiring that follows every exporter switch-case. Isolated from
// blockOtelInit so the parent selection stays flat and no individual case
// body exceeds Q4's 10-line PURE budget.
func otelTailLines() []string {
	return []string{
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
	}
}
