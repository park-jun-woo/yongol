//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=observability
//ff:what prometheusPath — metrics.path 결정 (미지정 시 "/metrics")
package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestPrometheusPath(t *testing.T) {
	cases := []struct {
		name string
		fs   *yongol.Fullstack
		want string
	}{
		{"nil fs", nil, "/metrics"},
		{"nil manifest", &yongol.Fullstack{}, "/metrics"},
		{"no metrics block", &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}, "/metrics"},
		{"empty path", fsWithMetrics(&pmanifest.ObservabilityMetrics{Path: ""}), "/metrics"},
		{"custom path", fsWithMetrics(&pmanifest.ObservabilityMetrics{Path: "/internal/metrics"}), "/internal/metrics"},
	}
	for _, c := range cases {
		if got := prometheusPath(c.fs); got != c.want {
			t.Errorf("%s: prometheusPath = %q, want %q", c.name, got, c.want)
		}
	}
}
