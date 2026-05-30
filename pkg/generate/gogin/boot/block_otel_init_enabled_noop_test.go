//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what TestBlockOtelInit_EnabledNoop — noop exporter 선택 시 exporter import 부재 검증

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockOtelInit_EnabledNoop(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Metadata: pmanifest.Metadata{Name: "zenflow"},
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				Observability: &pmanifest.Observability{
					Tracing: &pmanifest.ObservabilityTracing{
						Enabled:    true,
						Exporter:   "noop",
						SampleRate: 1.0,
					},
				},
			},
		},
	}
	block := blockOtelInit(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	imports := strings.Join(block.Imports, "\n")

	if !strings.Contains(body, "otel.SetTracerProvider") {
		t.Fatalf("missing SetTracerProvider, got:\n%s", body)
	}
	// noop exporter ships no exporter SDK import.
	if strings.Contains(imports, "otlptracegrpc") || strings.Contains(imports, "stdout/stdouttrace") {
		t.Fatalf("noop exporter must not pull exporter imports, got:\n%s", imports)
	}
	if !strings.Contains(body, "otelShutdown") {
		t.Fatalf("missing shutdown hook, got:\n%s", body)
	}
}
