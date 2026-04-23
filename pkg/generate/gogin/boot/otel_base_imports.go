//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what otelBaseImports — blockOtelInit 가 exporter 와 무관하게 항상 import 하는 모듈

package boot

// otelBaseImports returns the core OTel SDK imports required by every
// tracing branch. Exporter-specific imports are appended by
// otelOtlpImport / otelStdoutImport.
func otelBaseImports() []string {
	return []string{
		`"go.opentelemetry.io/otel"`,
		`"go.opentelemetry.io/otel/propagation"`,
		`sdktrace "go.opentelemetry.io/otel/sdk/trace"`,
		`"go.opentelemetry.io/otel/sdk/resource"`,
		`semconv "go.opentelemetry.io/otel/semconv/v1.26.0"`,
	}
}
