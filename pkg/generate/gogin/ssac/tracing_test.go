//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=observability
//ff:what tracingWrapCalls 단위 테스트 (tracing.enabled AND wrap_calls 둘 다일 때만 true)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func mkFS(tr *manifest.ObservabilityTracing) *yongol.Fullstack {
	mf := &manifest.ProjectConfig{}
	if tr != nil {
		mf.Backend.Observability = &manifest.Observability{Tracing: tr}
	}
	return &yongol.Fullstack{Manifest: mf}
}

func TestTracingWrapCalls(t *testing.T) {
	cases := []struct {
		name string
		fs   *yongol.Fullstack
		want bool
	}{
		{"nil fullstack", nil, false},
		{"nil manifest", &yongol.Fullstack{}, false},
		{"no observability", mkFS(nil), false},
		{"disabled tracing", mkFS(&manifest.ObservabilityTracing{Enabled: false, WrapCalls: true}), false},
		{"enabled but no wrap", mkFS(&manifest.ObservabilityTracing{Enabled: true, WrapCalls: false}), false},
		{"both on", mkFS(&manifest.ObservabilityTracing{Enabled: true, WrapCalls: true}), true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tracingWrapCalls(tc.fs); got != tc.want {
				t.Errorf("tracingWrapCalls = %v, want %v", got, tc.want)
			}
		})
	}
}
