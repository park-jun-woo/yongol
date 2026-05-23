//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-observability
//ff:what obs04TracingSampleRate — tracing.sample_rate가 0.0~1.0 범위인지 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestObs04TracingSampleRate(t *testing.T) {
	mk := func(enabled bool, rate float64) *yongol.Fullstack {
		return &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{
			Observability: &pm.Observability{Tracing: &pm.ObservabilityTracing{Enabled: enabled, SampleRate: rate}},
		}}}
	}
	cases := []TestObs04TracingSampleRateCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "nil_obs", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 0},
		{name: "nil_tracing", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Observability: &pm.Observability{}}}}, wantCount: 0},
		{name: "disabled", fs: mk(false, -1.0), wantCount: 0},
		{name: "zero", fs: mk(true, 0.0), wantCount: 0},
		{name: "one", fs: mk(true, 1.0), wantCount: 0},
		{name: "mid", fs: mk(true, 0.5), wantCount: 0},
		{name: "negative", fs: mk(true, -0.1), wantCount: 1},
		{name: "above_one", fs: mk(true, 1.5), wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runObs04TracingSampleRate(t, c)
		})
	}
}
