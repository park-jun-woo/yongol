//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what TestBlockOtelInit_EnabledOtlp — otlp exporter 선택 시 endpoint/sample_rate 전달 확인

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockOtelInit_EnabledOtlp(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Metadata: pmanifest.Metadata{Name: "zenflow"},
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				Observability: &pmanifest.Observability{
					Tracing: &pmanifest.ObservabilityTracing{
						Enabled:      true,
						Exporter:     "otlp",
						OtlpEndpoint: "collector.prod:4317",
						SampleRate:   0.1,
					},
				},
			},
		},
	}
	block := blockOtelInit(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	imports := strings.Join(block.Imports, "\n")
	if !strings.Contains(imports, "otlptracegrpc") {
		t.Fatalf("missing otlptracegrpc import, got:\n%s", imports)
	}
	if !strings.Contains(body, `"collector.prod:4317"`) {
		t.Fatalf("custom endpoint not propagated, got:\n%s", body)
	}
	if !strings.Contains(body, "otlpEndpoint") {
		t.Fatalf("expected otlpEndpoint local declared for otlp exporter, got:\n%s", body)
	}
	if !strings.Contains(body, "0.1") {
		t.Fatalf("sample_rate 0.1 not propagated, got:\n%s", body)
	}
}
