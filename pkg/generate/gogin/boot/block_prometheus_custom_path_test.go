//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what TestBlockPrometheus_CustomPath — manifest 에 /internal/metrics 지정 시 그대로 전달

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
