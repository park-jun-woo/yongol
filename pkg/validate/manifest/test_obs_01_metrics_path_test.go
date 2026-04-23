//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-observability
//ff:what OBS-001 테스트 — metrics.path golden (정상 경로 / 미설정)

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestObs01MetricsPath_Golden(t *testing.T) {
	cases := []struct {
		name     string
		fs       *yongol.Fullstack
	}{
		{
			name: "observability absent",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{}},
			},
		},
		{
			name: "path empty -> default",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Observability: &pmanifest.Observability{Metrics: &pmanifest.ObservabilityMetrics{}},
					},
				},
			},
		},
		{
			name: "leading slash path",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Observability: &pmanifest.Observability{
							Metrics: &pmanifest.ObservabilityMetrics{Path: "/internal/metrics"},
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := obs01MetricsPath(c.fs); len(got) != 0 {
				t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
			}
		})
	}
}
