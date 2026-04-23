//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what otelStdoutImport — stdouttrace exporter import 라인

package boot

// otelStdoutImport returns the single import for the stdouttrace exporter,
// added only when exporter == "stdout".
func otelStdoutImport() string {
	return `"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"`
}
