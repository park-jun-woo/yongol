//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what otelOtlpEndpoint — tracing.otlp_endpoint 값 결정 (미지정 시 "localhost:4317")

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

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
