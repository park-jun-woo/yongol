//ff:type feature=projectconfig type=model
//ff:what ObservabilityTracing — backend.observability.tracing 모델 (OTel 분산 추적 설정)

package manifest

// ObservabilityTracing controls OpenTelemetry trace instrumentation for the
// generated Go+Gin backend. The block is **opt-in** — when unset, tracing is
// disabled and the generated go.mod does not pull OTel dependencies so build
// size stays lean for projects that don't need distributed tracing.
//
//   - Enabled:      master switch. Missing or false → no OTel code emitted.
//   - ServiceName:  `service.name` resource attribute. Empty → metadata.name
//                   from manifest at codegen time.
//   - Exporter:     one of "otlp" | "stdout" | "noop". Validated by OBS-003.
//                   Default "noop" keeps SDK init paths exercised without
//                   talking to a collector.
//   - OtlpEndpoint: gRPC endpoint for the OTLP exporter ("host:port" form,
//                   no scheme). Only consulted when Exporter == "otlp".
//                   Default "localhost:4317".
//   - SampleRate:   head-based sampler ratio in [0.0, 1.0]. Validated by
//                   OBS-004. Default 1.0 (sample every trace).
//   - WrapCalls:    when true, yongol wraps every SSaC `@call` with an
//                   explicit `otel.Tracer("ssac").Start(...)` span. Off by
//                   default to avoid excessive span volume — flip on for
//                   deep debugging or low-traffic services.
//
// Env overrides (resolved in generated main.go):
//   BACKEND_OBSERVABILITY_TRACING_ENABLED
//   BACKEND_OBSERVABILITY_TRACING_SERVICE_NAME
//   BACKEND_OBSERVABILITY_TRACING_EXPORTER
//   BACKEND_OBSERVABILITY_TRACING_OTLP_ENDPOINT
//   BACKEND_OBSERVABILITY_TRACING_SAMPLE_RATE
type ObservabilityTracing struct {
	Enabled      bool    `yaml:"enabled,omitempty"`
	ServiceName  string  `yaml:"service_name,omitempty"`
	Exporter     string  `yaml:"exporter,omitempty"`
	OtlpEndpoint string  `yaml:"otlp_endpoint,omitempty"`
	SampleRate   float64 `yaml:"sample_rate,omitempty"`
	WrapCalls    bool    `yaml:"wrap_calls,omitempty"`
}
