//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what C-6 테스트 — backend.auth 존재 시 진단 0 건 (golden)

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestC06BackendAuthRequired_AuthPresent_NoDiag — 정상 케이스: auth 블록이
// 존재하면 진단 0 건.
func TestC06BackendAuthRequired_AuthPresent_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "github.com/park-jun-woo/zenflow",
				Auth: &pmanifest.Auth{
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
