//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what otelOtlpImport — OTLP gRPC exporter import 라인

package boot

// otelOtlpImport returns the single import for the OTLP gRPC exporter,
// added only when exporter == "otlp".
func otelOtlpImport() string {
	return `otlpgrpc "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"`
}
