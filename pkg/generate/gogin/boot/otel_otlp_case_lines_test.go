//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=observability
//ff:what otelOtlpCaseLines — switch "otlp" case 의 본문 라인 생성 (endpoint 주입)
package boot

import (
	"strings"
	"testing"
)

func TestOtelOtlpCaseLines(t *testing.T) {
	lines := otelOtlpCaseLines("collector:4317")
	body := strings.Join(lines, "\n")
	for _, must := range []string{
		`case "otlp":`,
		`envString("BACKEND_OBSERVABILITY_TRACING_OTLP_ENDPOINT", "collector:4317")`,
		"otlpgrpc.New(ctx,",
		"otlpgrpc.WithInsecure(),",
		"spanExporter = exp",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("otlp branch missing %q, got:\n%s", must, body)
		}
	}
}
