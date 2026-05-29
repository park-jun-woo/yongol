//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what otelExporter — tracing.exporter 값 결정 (미지정 시 "noop")

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestOtelExporter(t *testing.T) {
	if got := otelExporter(nil); got != "noop" {
		t.Errorf("no otel = %q, want noop", got)
	}
	empty := fsTracing(&pmanifest.ObservabilityTracing{Enabled: true, Exporter: ""})
	if got := otelExporter(empty); got != "noop" {
		t.Errorf("empty exporter = %q, want noop", got)
	}
	otlp := fsTracing(&pmanifest.ObservabilityTracing{Enabled: true, Exporter: "otlp"})
	if got := otelExporter(otlp); got != "otlp" {
		t.Errorf("explicit exporter = %q, want otlp", got)
	}
}
