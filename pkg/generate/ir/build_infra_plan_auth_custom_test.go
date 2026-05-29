//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildInfraPlanAuthCustom -- auth 커스텀 TTL/SecretEnv 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildInfraPlanAuthCustom(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Module: "github.com/test/auth-custom",
				Auth: &manifest.Auth{
					SecretEnv:       "MY_JWT_KEY",
					AccessTokenTTL:  "30m",
					RefreshTokenTTL: "720h",
				},
			},
		},
	}
	ps := &prepared.State{
		Auth: prepared.Auth{
			Present: true,
			Mode:    "cookie",
			Raw:     fs.Manifest.Backend.Auth,
		},
	}

	plan := BuildInfraPlan(fs, ps)

	if plan.Auth == nil {
		t.Fatal("Auth should be non-nil")
	}
	if plan.Auth.SecretEnv != "MY_JWT_KEY" {
		t.Errorf("Auth.SecretEnv = %q, want %q", plan.Auth.SecretEnv, "MY_JWT_KEY")
	}
	if plan.Auth.AccessTokenTTL != "30m" {
		t.Errorf("Auth.AccessTokenTTL = %q, want %q", plan.Auth.AccessTokenTTL, "30m")
	}
	if plan.Auth.RefreshTokenTTL != "720h" {
		t.Errorf("Auth.RefreshTokenTTL = %q, want %q", plan.Auth.RefreshTokenTTL, "720h")
	}
}
