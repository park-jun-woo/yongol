//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what otelOtlpImport — OTLP gRPC exporter import 라인

package boot

import (
	"strings"
	"testing"
)

func TestOtelOtlpImport(t *testing.T) {
	imp := otelOtlpImport()
	if !strings.HasPrefix(imp, "otlpgrpc ") {
		t.Errorf("otlp import should be aliased otlpgrpc, got %q", imp)
	}
	if !strings.Contains(imp, "otlptracegrpc") {
		t.Errorf("otlp import should target otlptracegrpc, got %q", imp)
	}
}
