//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what queueCfg — manifest.queue 섹션을 builtinBackend로 래핑 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQueueCfg(t *testing.T) {
	cases := []struct {
		name        string
		fs          *yongol.Fullstack
		wantPresent bool
		wantBackend string
	}{
		{name: "nil_fs", fs: nil, wantPresent: false},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantPresent: false},
		{name: "nil_queue", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantPresent: false},
		{name: "present", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Queue: &pm.QueueBackend{Backend: "postgres"}}}, wantPresent: true, wantBackend: "postgres"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := queueCfg(c.fs)
			if got.Present != c.wantPresent {
				t.Errorf("Present = %v, want %v", got.Present, c.wantPresent)
			}
			if got.Backend != c.wantBackend {
				t.Errorf("Backend = %q, want %q", got.Backend, c.wantBackend)
			}
		})
	}
}
