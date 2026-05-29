//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildMiddlewarePlanBearerAuth -- bearer 인증 → BearerAuth 설정 + CSRF nil 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildMiddlewarePlanBearerAuth(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Module: "github.com/test/auth",
				Auth: &manifest.Auth{
					SecretEnv: "MY_SECRET",
					Claims:    map[string]manifest.ClaimDef{"UserID": {Key: "uid"}},
				},
			},
		},
	}
	ps := &prepared.State{
		Auth: prepared.Auth{
			Present: true,
			Mode:    "bearer",
			Raw:     fs.Manifest.Backend.Auth,
		},
	}

	plan := BuildMiddlewarePlan(fs, ps)

	if plan.BearerAuth == nil {
		t.Fatal("BearerAuth should be non-nil when auth is present")
	}
	if plan.BearerAuth.Mode != "bearer" {
		t.Errorf("BearerAuth.Mode = %q, want %q", plan.BearerAuth.Mode, "bearer")
	}
	if plan.BearerAuth.SecretEnv != "MY_SECRET" {
		t.Errorf("BearerAuth.SecretEnv = %q, want %q", plan.BearerAuth.SecretEnv, "MY_SECRET")
	}
	if !plan.BearerAuth.HasClaims {
		t.Error("BearerAuth.HasClaims should be true")
	}
	if plan.CSRF != nil {
		t.Error("CSRF should be nil for bearer-only auth")
	}
}
