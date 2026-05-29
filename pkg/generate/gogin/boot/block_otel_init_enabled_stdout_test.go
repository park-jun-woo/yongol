//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what TestBlockOtelInit_EnabledStdout — stdout exporter 선택 시 심볼/import 검증

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockOtelInit_EnabledStdout(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Metadata: pmanifest.Metadata{Name: "zenflow"},
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				Observability: &pmanifest.Observability{
					Tracing: &pmanifest.ObservabilityTracing{
						Enabled:    true,
						Exporter:   "stdout",
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
	if !strings.Contains(body, "propagation.TraceContext{}") {
		t.Fatalf("missing W3C TraceContext propagator, got:\n%s", body)
	}
	if !strings.Contains(body, "stdouttrace.New") {
		t.Fatalf("missing stdout exporter branch, got:\n%s", body)
	}
	if !strings.Contains(imports, "stdout/stdouttrace") {
		t.Fatalf("missing stdouttrace import, got:\n%s", imports)
	}
	if strings.Contains(imports, "otlptracegrpc") {
		t.Fatalf("otlp exporter import should not be present for stdout exporter, got:\n%s", imports)
	}
	if !strings.Contains(body, "otelShutdown") {
		t.Fatalf("missing shutdown hook, got:\n%s", body)
	}
	// service.name defaults to metadata.name when tracing.service_name is unset
	if !strings.Contains(body, `"zenflow"`) {
		t.Fatalf("expected service.name defaulted to metadata.name, got:\n%s", body)
	}
}
