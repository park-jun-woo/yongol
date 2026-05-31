//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConfigBuildersZeroCov — boot/middleware config 빌더 전 분기 직접 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBuildBearerAuthConfig_ZeroCov(t *testing.T) {
	// not present → nil.
	if c := buildBearerAuthConfig(&prepared.State{}); c != nil {
		t.Errorf("absent auth → nil")
	}
	// present, custom secret + claims.
	ps := &prepared.State{Auth: prepared.Auth{
		Present: true, Mode: "bearer",
		Raw: &manifest.Auth{SecretEnv: "MY_SECRET", Claims: map[string]manifest.ClaimDef{"sub": {}}},
	}}
	c := buildBearerAuthConfig(ps)
	if c == nil || c.SecretEnv != "MY_SECRET" || !c.HasClaims {
		t.Errorf("bearer auth = %+v", c)
	}
	// present, no raw → default secret env.
	ps2 := &prepared.State{Auth: prepared.Auth{Present: true, Mode: "bearer"}}
	if c := buildBearerAuthConfig(ps2); c.SecretEnv != "JWT_SECRET" {
		t.Errorf("default secret env = %q", c.SecretEnv)
	}
}
