//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-observability
//ff:what resolveMetricsPath — metrics.path 유효 값 결정 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveMetricsPath(t *testing.T) {
	mk := func(path string) *yongol.Fullstack {
		return &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{
			Observability: &pm.Observability{Metrics: &pm.ObservabilityMetrics{Path: path}},
		}}}
	}
	cases := []struct {
		name string
		fs   *yongol.Fullstack
		want string
	}{
		{name: "nil_obs", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, want: "/metrics"},
		{name: "nil_metrics", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Observability: &pm.Observability{}}}}, want: "/metrics"},
		{name: "empty_path", fs: mk(""), want: "/metrics"},
		{name: "valid_path", fs: mk("/custom"), want: "/custom"},
		{name: "no_leading_slash", fs: mk("metrics"), want: ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := resolveMetricsPath(c.fs)
			if got != c.want {
				t.Errorf("resolveMetricsPath() = %q, want %q", got, c.want)
			}
		})
	}
}
