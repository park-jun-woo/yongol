//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what otelOtlpEndpoint — tracing.otlp_endpoint 값 결정 (미지정 시 "localhost:4317")

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestOtelOtlpEndpoint(t *testing.T) {
	if got := otelOtlpEndpoint(nil); got != "localhost:4317" {
		t.Errorf("no otel = %q, want localhost:4317", got)
	}
	empty := fsTracing(&pmanifest.ObservabilityTracing{Enabled: true, OtlpEndpoint: ""})
	if got := otelOtlpEndpoint(empty); got != "localhost:4317" {
		t.Errorf("empty endpoint = %q, want localhost:4317", got)
	}
	set := fsTracing(&pmanifest.ObservabilityTracing{Enabled: true, OtlpEndpoint: "collector:4317"})
	if got := otelOtlpEndpoint(set); got != "collector:4317" {
		t.Errorf("explicit endpoint = %q, want collector:4317", got)
	}
}
