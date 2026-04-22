//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what hasOtel — manifest.backend.observability.tracing.enabled 여부 (기본 false)

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasOtel returns true when OpenTelemetry tracing is explicitly enabled in
// manifest.backend.observability.tracing. Defaults to **false** (opt-in):
// a missing observability block, missing tracing block, or tracing.enabled:
// false all yield false so projects that don't need distributed tracing
// incur zero OTel build / runtime cost.
func hasOtel(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil {
		return false
	}
	obs := fs.Manifest.Backend.Observability
	if obs == nil || obs.Tracing == nil {
		return false
	}
	return obs.Tracing.Enabled
}

// otelWrapCalls reports whether `@call` sites should be wrapped with an
// explicit ssac-tracer span. Requires hasOtel (master toggle) AND the
// wrap_calls opt-in so ordinary services don't get a span per call.
func otelWrapCalls(fs *yongol.Fullstack) bool {
	if !hasOtel(fs) {
		return false
	}
	return fs.Manifest.Backend.Observability.Tracing.WrapCalls
}

// otelServiceName resolves the `service.name` resource attribute: manifest
// tracing.service_name when set, else metadata.name (the project name), else
// "service" as a last-resort fallback so the exporter never receives an
// empty identifier.
func otelServiceName(fs *yongol.Fullstack) string {
	if fs == nil || fs.Manifest == nil {
		return "service"
	}
	obs := fs.Manifest.Backend.Observability
	if obs != nil && obs.Tracing != nil && obs.Tracing.ServiceName != "" {
		return obs.Tracing.ServiceName
	}
	if fs.Manifest.Metadata.Name != "" {
		return fs.Manifest.Metadata.Name
	}
	return "service"
}

// otelExporter resolves the exporter kind with "noop" as the safe default
// when the field is unset. Validation (OBS-003) restricts explicit values
// to otlp/stdout/noop so unknown strings never reach here.
func otelExporter(fs *yongol.Fullstack) string {
	if !hasOtel(fs) {
		return "noop"
	}
	v := fs.Manifest.Backend.Observability.Tracing.Exporter
	if v == "" {
		return "noop"
	}
	return v
}

// otelOtlpEndpoint returns the gRPC endpoint for the OTLP exporter, falling
// back to the canonical dev-default "localhost:4317" (Jaeger all-in-one,
// otel-collector) when unset.
func otelOtlpEndpoint(fs *yongol.Fullstack) string {
	if !hasOtel(fs) {
		return "localhost:4317"
	}
	v := fs.Manifest.Backend.Observability.Tracing.OtlpEndpoint
	if v == "" {
		return "localhost:4317"
	}
	return v
}

// otelSampleRate returns the head-sampler ratio clamped to a default of 1.0
// when the field is zero AND tracing was only implicitly enabled (since
// Go's zero value for float64 is 0.0, indistinguishable from explicit "0").
// To preserve the "explicit 0 = never sample" semantic, the helper treats
// **any** non-positive value as the 1.0 default so users who truly want 0
// must be explicit via env override at runtime.
func otelSampleRate(fs *yongol.Fullstack) float64 {
	if !hasOtel(fs) {
		return 1.0
	}
	r := fs.Manifest.Backend.Observability.Tracing.SampleRate
	if r <= 0 {
		return 1.0
	}
	return r
}
