//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what otelOtlpCaseLines — switch "otlp" case 의 본문 라인 생성 (endpoint 주입)

package boot

import "fmt"

// otelOtlpCaseLines emits the `case "otlp":` branch body. endpoint is the
// default for the BACKEND_OBSERVABILITY_TRACING_OTLP_ENDPOINT env var.
func otelOtlpCaseLines(endpoint string) []string {
	return []string{
		`	case "otlp":`,
		fmt.Sprintf(`		otlpEndpoint := envString("BACKEND_OBSERVABILITY_TRACING_OTLP_ENDPOINT", %q)`, endpoint),
		`		exp, err := otlpgrpc.New(ctx,`,
		`			otlpgrpc.WithEndpoint(otlpEndpoint),`,
		`			otlpgrpc.WithInsecure(),`,
		`		)`,
		`		if err != nil { slog.Error("otel otlp exporter", "err", err); os.Exit(1) }`,
		`		spanExporter = exp`,
	}
}
