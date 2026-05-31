//ff:func feature=gen-gogin type=test control=sequence
//ff:what resolveAuthInitConfig — prepared.Auth 에서 blockAuthInit 가 필요한 모든 값 해석
package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestResolveAuthInitConfig_Overrides(t *testing.T) {
	raw := &pmanifest.Auth{
		SecretEnv:            "MY_SECRET",
		AccessTokenTTL:       "5m",
		RefreshTokenTTL:      "24h",
		DetectReuseLogoutAll: true,
		Cookie: &pmanifest.CookieConfig{
			SameSite:    "Strict",
			AccessName:  "ac",
			RefreshName: "rf",
		},
	}
	a := prepared.Auth{Present: true, Mode: "hybrid", Raw: raw}
	cfg := resolveAuthInitConfig(a)
	if cfg.SecretEnv != "MY_SECRET" || cfg.AccessTTL != "5m" || cfg.RefreshTTL != "24h" {
		t.Errorf("token overrides not applied: %+v", cfg)
	}
	if cfg.Mode != "hybrid" {
		t.Errorf("Mode should come from prepared.Auth.Mode, got %q", cfg.Mode)
	}
	if !cfg.DetectReuse {
		t.Errorf("DetectReuse should be true")
	}
	if cfg.SameSite != "Strict" || cfg.AccessName != "ac" || cfg.RefreshName != "rf" {
		t.Errorf("cookie overrides not applied: %+v", cfg)
	}
}
