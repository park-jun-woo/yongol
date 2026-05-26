//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildInfraPlanAuthDefaults -- auth 기본값 (SecretEnv/TTL) 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildInfraPlanAuthDefaults(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Module: "github.com/test/auth",
				Auth:   &manifest.Auth{},
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

	plan := BuildInfraPlan(fs, ps)

	if plan.Auth == nil {
		t.Fatal("Auth should be non-nil when auth is present")
	}
	if plan.Auth.Mode != "bearer" {
		t.Errorf("Auth.Mode = %q, want %q", plan.Auth.Mode, "bearer")
	}
	if plan.Auth.SecretEnv != "JWT_SECRET" {
		t.Errorf("Auth.SecretEnv = %q, want %q", plan.Auth.SecretEnv, "JWT_SECRET")
	}
	if plan.Auth.AccessTokenTTL != "15m" {
		t.Errorf("Auth.AccessTokenTTL = %q, want %q", plan.Auth.AccessTokenTTL, "15m")
	}
	if plan.Auth.RefreshTokenTTL != "168h" {
		t.Errorf("Auth.RefreshTokenTTL = %q, want %q", plan.Auth.RefreshTokenTTL, "168h")
	}
}
