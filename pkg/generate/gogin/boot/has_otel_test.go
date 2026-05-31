//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=observability
//ff:what hasOtel — manifest.backend.observability.tracing.enabled 여부 (기본 false)
package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestHasOtel(t *testing.T) {
	cases := []struct {
		name string
		fs   *yongol.Fullstack
		want bool
	}{
		{"nil fs", nil, false},
		{"nil manifest", &yongol.Fullstack{}, false},
		{"no observability block", &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}, false},
		{"tracing nil", fsWithTracing(nil), false},
		{"tracing disabled", fsWithTracing(&pmanifest.ObservabilityTracing{Enabled: false}), false},
		{"tracing enabled", fsWithTracing(&pmanifest.ObservabilityTracing{Enabled: true}), true},
	}
	for _, c := range cases {
		if got := hasOtel(c.fs); got != c.want {
			t.Errorf("%s: hasOtel = %v, want %v", c.name, got, c.want)
		}
	}
}
