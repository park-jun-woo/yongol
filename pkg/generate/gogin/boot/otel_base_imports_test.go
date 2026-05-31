//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=observability
//ff:what otelBaseImports — blockOtelInit 가 exporter 와 무관하게 항상 import 하는 모듈
package boot

import (
	"strings"
	"testing"
)

func TestOtelBaseImports(t *testing.T) {
	imps := otelBaseImports()
	if len(imps) == 0 {
		t.Fatalf("expected base imports, got none")
	}
	joined := strings.Join(imps, "\n")
	for _, want := range []string{
		`"go.opentelemetry.io/otel"`,
		`"go.opentelemetry.io/otel/propagation"`,
		`sdktrace "go.opentelemetry.io/otel/sdk/trace"`,
		`"go.opentelemetry.io/otel/sdk/resource"`,
		`semconv "go.opentelemetry.io/otel/semconv/v1.26.0"`,
	} {
		if !strings.Contains(joined, want) {
			t.Errorf("missing base import %q in:\n%s", want, joined)
		}
	}
}
