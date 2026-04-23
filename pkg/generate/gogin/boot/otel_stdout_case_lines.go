//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what otelStdoutCaseLines — switch "stdout" case 의 본문 라인 생성

package boot

// otelStdoutCaseLines emits the `case "stdout":` branch body. Pretty-print
// is enabled so dev logs are human-readable; keep production on OTLP.
func otelStdoutCaseLines() []string {
	return []string{
		`	case "stdout":`,
		`		exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())`,
		`		if err != nil { slog.Error("otel stdout exporter", "err", err); os.Exit(1) }`,
		`		spanExporter = exp`,
	}
}
