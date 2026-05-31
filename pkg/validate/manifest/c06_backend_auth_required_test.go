//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what c06BackendAuthRequired — manifest backend.auth 블록 필수 검증
package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC06BackendAuthRequired(t *testing.T) {
	cases := []TestC06BackendAuthRequiredCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "auth_present", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{}}}}, wantCount: 0},
		{name: "auth_nil", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runC06BackendAuthRequired(t, c)
		})
	}
}
