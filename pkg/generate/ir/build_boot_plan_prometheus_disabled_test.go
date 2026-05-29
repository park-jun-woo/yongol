//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestBuildBootPlanPrometheusDisabled -- prometheus 명시 비활성화 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildBootPlanPrometheusDisabled(t *testing.T) {
	disabled := false
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Metadata: manifest.Metadata{Name: "nopromproj"},
			Backend: manifest.Backend{
				Module: "github.com/test/noprom",
				Observability: &manifest.Observability{
					Metrics: &manifest.ObservabilityMetrics{Enabled: &disabled},
				},
			},
		},
	}
	ps := &prepared.State{}

	plan := BuildBootPlan(fs, ps, "go-gin")

	blockMap := map[string]bool{}
	for _, b := range plan.ActiveBlocks {
		blockMap[b.Name] = b.Active
	}

	if blockMap["prometheus"] {
		t.Error("prometheus should be inactive when explicitly disabled")
	}
}
