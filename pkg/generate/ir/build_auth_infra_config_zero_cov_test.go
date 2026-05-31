//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConfigBuildersZeroCov — boot/middleware config 빌더 전 분기 직접 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildAuthInfraConfig_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	// no raw → all defaults.
	ps := &prepared.State{Auth: prepared.Auth{Present: true, Mode: "cookie"}}
	c := buildAuthInfraConfig(fs, ps)
	if c.SecretEnv != "JWT_SECRET" || c.AccessTokenTTL != "15m" || c.RefreshTokenTTL != "168h" {
		t.Errorf("auth infra defaults = %+v", c)
	}
	// raw with custom values → overrides.
	ps2 := &prepared.State{Auth: prepared.Auth{
		Present: true, Mode: "cookie",
		Raw: &manifest.Auth{SecretEnv: "S", AccessTokenTTL: "5m", RefreshTokenTTL: "24h"},
	}}
	c2 := buildAuthInfraConfig(fs, ps2)
	if c2.SecretEnv != "S" || c2.AccessTokenTTL != "5m" || c2.RefreshTokenTTL != "24h" {
		t.Errorf("auth infra overrides = %+v", c2)
	}
}
