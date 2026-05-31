//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-observability
//ff:what obs01MetricsPath — metrics.path가 /로 시작하는지 검증
package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestObs01MetricsPath(t *testing.T) {
	cases := []TestObs01MetricsPathCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "nil_observability", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 0},
		{name: "nil_metrics", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Observability: &pm.Observability{}}}}, wantCount: 0},
		{name: "empty_path_ok", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Observability: &pm.Observability{Metrics: &pm.ObservabilityMetrics{Path: ""}}}}}, wantCount: 0},
		{name: "valid_path", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Observability: &pm.Observability{Metrics: &pm.ObservabilityMetrics{Path: "/metrics"}}}}}, wantCount: 0},
		{name: "invalid_path", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Observability: &pm.Observability{Metrics: &pm.ObservabilityMetrics{Path: "metrics"}}}}}, wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runObs01MetricsPath(t, c)
		})
	}
}
