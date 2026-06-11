//ff:func feature=gen-react type=test control=sequence
//ff:what generateFrontendSetup — cookie 모드 + claims 캡처 fixture 에서 claims 전용 store 방출, 캡처 부재 cookie 프로젝트는 store 미방출 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateFrontendSetup_CookieClaims(t *testing.T) {
	makeCookieFS := func(pages []stml.PageSpec) *yongol.Fullstack {
		return &yongol.Fullstack{
			STMLPages: pages,
			Manifest: &manifest.ProjectConfig{
				Backend: manifest.Backend{Auth: &manifest.Auth{Type: "session", Mode: "cookie", Roles: []string{"member", "admin"}}},
				Frontend: manifest.Frontend{
					Lang: "typescript",
					Auth: &manifest.FrontendAuth{RoleField: "role"},
				},
			},
		}
	}
	claimsPage := stml.PageSpec{Name: "login", FileName: "login.html", Actions: []stml.ActionBlock{{
		OperationID: "Login",
		CaptureRaw:  "role -> auth.claims.role",
		Captures:    []stml.CaptureBind{{RespField: "role", Sink: "auth.claims.role"}},
	}}}

	t.Run("cookie mode with a claims capture emits the claims-only store", func(t *testing.T) {
		out := t.TempDir()
		if err := generateFrontendSetup(makeCookieFS([]stml.PageSpec{claimsPage}), out); err != nil {
			t.Fatalf("generateFrontendSetup: %v", err)
		}
		data, err := os.ReadFile(filepath.Join(out, "frontend", "src", "stores", "auth.ts"))
		if err != nil {
			t.Fatalf("expected stores/auth.ts in cookie mode with claims capture: %v", err)
		}
		store := string(data)
		assertContains(t, store, "claims: Record<string, string>")
		assertContains(t, store, "setClaim")
		assertContains(t, store, "clear: () => set({ claims: {} }),")
		// Reduced shape: no token surface in cookie mode.
		assertNotContains(t, store, "setAuth")
		assertNotContains(t, store, "token")
	})

	t.Run("cookie mode without claims captures emits no store (pre-Phase005 shape)", func(t *testing.T) {
		out := t.TempDir()
		plain := stml.PageSpec{Name: "login", FileName: "login.html", Actions: []stml.ActionBlock{{OperationID: "Login"}}}
		if err := generateFrontendSetup(makeCookieFS([]stml.PageSpec{plain}), out); err != nil {
			t.Fatalf("generateFrontendSetup: %v", err)
		}
		if _, err := os.Stat(filepath.Join(out, "frontend", "src", "stores", "auth.ts")); !os.IsNotExist(err) {
			t.Errorf("cookie mode without claims must not emit stores/auth.ts (err=%v)", err)
		}
	})
}
