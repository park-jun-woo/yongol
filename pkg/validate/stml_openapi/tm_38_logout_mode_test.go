//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-38 — auth 부재/cookie+값 없음 발화, bearer·cookie+op·jwt 무모드(bearer 유도) 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestTM38LogoutMode(t *testing.T) {
	fsWith := func(auth *manifest.Auth, logout *stml.LogoutSpec) *yongol.Fullstack {
		fs := makeFS(nil, nil)
		fs.Manifest.Backend.Auth = auth
		fs.Layouts = []stml.LayoutSpec{{Name: "app", File: "layouts/app.html", Logout: logout}}
		return fs
	}

	t.Run("no auth with logout is a warning", func(t *testing.T) {
		diags := tm38LogoutMode(fsWith(nil, &stml.LogoutSpec{OperationID: "Logout"}))
		if got := countDiag(diags, "[TM-38]"); got != 1 {
			t.Fatalf("expected 1 TM-38, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelWarning {
			t.Errorf("Level = %v, want LevelWarning", diags[0].Level)
		}
		if !strings.Contains(diags[0].Message, "no backend.auth") {
			t.Errorf("Message = %q, want no-auth branch", diags[0].Message)
		}
	})

	t.Run("cookie mode valueless logout is a warning", func(t *testing.T) {
		diags := tm38LogoutMode(fsWith(&manifest.Auth{Type: "jwt", Mode: "cookie"}, &stml.LogoutSpec{}))
		if got := countDiag(diags, "[TM-38]"); got != 1 {
			t.Fatalf("expected 1 TM-38, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "httpOnly cookie") {
			t.Errorf("Message = %q, want cookie branch", diags[0].Message)
		}
	})

	t.Run("cookie mode with op is silent", func(t *testing.T) {
		diags := tm38LogoutMode(fsWith(&manifest.Auth{Type: "jwt", Mode: "cookie"}, &stml.LogoutSpec{OperationID: "Logout"}))
		if len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("bearer mode valueless logout is silent", func(t *testing.T) {
		diags := tm38LogoutMode(fsWith(&manifest.Auth{Type: "jwt", Mode: "bearer"}, &stml.LogoutSpec{}))
		if len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("jwt without mode derives bearer and is silent", func(t *testing.T) {
		// prepared.AuthFor maps jwt-without-mode to bearer (BUG-014) — the
		// emitter emits a store-clearing logout, so no warning may fire.
		diags := tm38LogoutMode(fsWith(&manifest.Auth{Type: "jwt"}, &stml.LogoutSpec{}))
		if len(diags) != 0 {
			t.Errorf("expected silence for jwt-derived bearer, got %+v", diags)
		}
	})

	t.Run("no logout is silent", func(t *testing.T) {
		diags := tm38LogoutMode(fsWith(nil, nil))
		if len(diags) != 0 {
			t.Errorf("expected silence without data-logout, got %+v", diags)
		}
	})
}
