//ff:func feature=gen-gogin type=test control=sequence
//ff:what resolveAuthInitConfig — prepared.Auth 에서 blockAuthInit 가 필요한 모든 값 해석
package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

func TestResolveAuthInitConfig_Defaults(t *testing.T) {
	// Not present / nil raw → full default bundle.
	cfg := resolveAuthInitConfig(prepared.Auth{Present: false})
	if cfg.SecretEnv != "JWT_SECRET" || cfg.AccessTTL != "15m" || cfg.RefreshTTL != "168h" {
		t.Errorf("default token cfg wrong: %+v", cfg)
	}
	if cfg.Mode != "cookie" || cfg.SameSite != "Lax" {
		t.Errorf("default mode/samesite wrong: %+v", cfg)
	}
	if cfg.AccessName != "__Host-access_token" || cfg.RefreshName != "__Host-refresh_token" {
		t.Errorf("default cookie names wrong: %+v", cfg)
	}
}
