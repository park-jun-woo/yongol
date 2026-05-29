//ff:func feature=gen-gogin type=test control=sequence
//ff:what resolveAuthCookieConfig — manifest.auth.cookie 서브블록 값으로 authInitConfig 덮어쓰기

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestResolveAuthCookieConfig(t *testing.T) {
	base := func() authInitConfig {
		return authInitConfig{SameSite: "Lax", AccessName: "a", RefreshName: "r"}
	}

	t.Run("nil cookie leaves cfg untouched", func(t *testing.T) {
		cfg := base()
		resolveAuthCookieConfig(nil, &cfg)
		if cfg.SameSite != "Lax" || cfg.AccessName != "a" || cfg.RefreshName != "r" {
			t.Errorf("nil cookie mutated cfg: %+v", cfg)
		}
	})

	t.Run("empty fields skip override", func(t *testing.T) {
		cfg := base()
		resolveAuthCookieConfig(&pmanifest.CookieConfig{}, &cfg)
		if cfg.SameSite != "Lax" || cfg.AccessName != "a" || cfg.RefreshName != "r" {
			t.Errorf("empty cookie fields overrode cfg: %+v", cfg)
		}
	})

	t.Run("set fields override", func(t *testing.T) {
		cfg := base()
		resolveAuthCookieConfig(&pmanifest.CookieConfig{SameSite: "None", AccessName: "x", RefreshName: "y"}, &cfg)
		if cfg.SameSite != "None" || cfg.AccessName != "x" || cfg.RefreshName != "y" {
			t.Errorf("overrides not applied: %+v", cfg)
		}
	})
}
