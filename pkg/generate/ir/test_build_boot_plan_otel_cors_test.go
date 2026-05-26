//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestBuildBootPlanOtelCors -- otel-init + cors 블록 조건부 활성화 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildBootPlanOtelCors(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Metadata: manifest.Metadata{Name: "otelcorsproj"},
			Backend: manifest.Backend{
				Module: "github.com/test/otelcors",
				Observability: &manifest.Observability{
					Tracing: &manifest.ObservabilityTracing{Enabled: true},
				},
				CORS: &manifest.CORSConfig{Enabled: true},
			},
		},
	}
	ps := &prepared.State{}

	plan := BuildBootPlan(fs, ps, "go-gin")

	blockMap := map[string]bool{}
	for _, b := range plan.ActiveBlocks {
		blockMap[b.Name] = b.Active
	}

	if !blockMap["otel-init"] {
		t.Error("otel-init should be active when tracing is enabled")
	}
	if !blockMap["cors"] {
		t.Error("cors should be active when manifest.backend.cors.enabled is true")
	}
}
