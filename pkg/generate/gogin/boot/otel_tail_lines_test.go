//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what otelTailLines — exporter 분기 이후 공통 TracerProvider 초기화 + shutdown

package boot

import (
	"strings"
	"testing"
)

func TestOtelTailLines(t *testing.T) {
	body := strings.Join(otelTailLines(), "\n")
	for _, must := range []string{
		"default:",
		"resource.New(ctx,",
		"sdktrace.NewTracerProvider(tpOpts...)",
		"otel.SetTracerProvider(tp)",
		"otel.SetTextMapPropagator(",
		"otelShutdown = tp.Shutdown",
		"defer func() {",
		"otelShutdown(shutdownCtx)",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("tail lines missing %q, got:\n%s", must, body)
		}
	}
}
