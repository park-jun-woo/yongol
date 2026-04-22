//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what blockPrometheus 활성/비활성 + 경로 override 테스트

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockPrometheus_DefaultActive(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Module: "example.com/zenflow"},
		},
	}
	block := blockPrometheus(fs, "example.com/zenflow")
	if block.Name != "prometheus" {
		t.Fatalf("unexpected name %q", block.Name)
	}
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "middleware.PrometheusMiddleware()") {
		t.Fatalf("expected PrometheusMiddleware registration, got:\n%s", body)
	}
	if !strings.Contains(body, "middleware.PrometheusHandler()") {
		t.Fatalf("expected PrometheusHandler registration, got:\n%s", body)
	}
	if !strings.Contains(body, `"/metrics"`) {
		t.Fatalf("expected default /metrics path, got:\n%s", body)
	}
	if !strings.Contains(body, `BACKEND_OBSERVABILITY_METRICS_ENABLED`) {
		t.Fatalf("expected env toggle, got:\n%s", body)
	}
}

func TestBlockPrometheus_DisabledExplicitly(t *testing.T) {
	disabled := false
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				Observability: &pmanifest.Observability{
					Metrics: &pmanifest.ObservabilityMetrics{Enabled: &disabled},
				},
			},
		},
	}
	block := blockPrometheus(fs, "example.com/zenflow")
	if len(block.Lines) != 0 {
		t.Fatalf("expected inert block when metrics disabled, got lines: %+v", block.Lines)
	}
}

func TestBlockPrometheus_CustomPath(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				Observability: &pmanifest.Observability{
					Metrics: &pmanifest.ObservabilityMetrics{Path: "/internal/metrics"},
				},
			},
		},
	}
	block := blockPrometheus(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `"/internal/metrics"`) {
		t.Fatalf("expected custom path in block, got:\n%s", body)
	}
}
