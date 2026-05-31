//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what c06BackendAuthRequired — manifest backend.auth 블록 필수 검증
package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
