//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what TestBlockPrometheus_DefaultActive — 기본 설정에서 /metrics + middleware 등록

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
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
