//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what sessionCfg — manifest.session 섹션을 builtinBackend로 래핑 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSessionCfg(t *testing.T) {
	cases := []struct {
		name        string
		fs          *yongol.Fullstack
		wantPresent bool
		wantBackend string
	}{
		{name: "nil_fs", fs: nil, wantPresent: false},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantPresent: false},
		{name: "nil_session", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantPresent: false},
		{name: "present", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Session: &pm.BuiltinBackend{Backend: "memory"}}}, wantPresent: true, wantBackend: "memory"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := sessionCfg(c.fs)
			if got.Present != c.wantPresent {
				t.Errorf("Present = %v, want %v", got.Present, c.wantPresent)
			}
			if got.Backend != c.wantBackend {
				t.Errorf("Backend = %q, want %q", got.Backend, c.wantBackend)
			}
		})
	}
}
