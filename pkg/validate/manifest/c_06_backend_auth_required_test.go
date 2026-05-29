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

// TestC06BackendAuthRequired_AuthPresent_NoDiag — 정상 케이스: auth 블록이
// 존재하면 진단 0 건.
func TestC06BackendAuthRequired_AuthPresent_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Module: "github.com/park-jun-woo/zenflow",
				Auth: &pm.Auth{
					Type:      "jwt",
					SecretEnv: "JWT_SECRET",
					RawClaims: map[string]string{
						"ID":   "user_id:string",
						"Role": "role",
					},
				},
			},
		},
	}
	if got := c06BackendAuthRequired(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
