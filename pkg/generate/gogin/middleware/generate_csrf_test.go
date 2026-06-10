//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what TestGenerateCsrf — csrf.go 기록/스킵/모드 주입/에러 경로 검증
package middleware

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestGenerateCsrf(t *testing.T) {
	t.Run("SkipsWhenAuthAbsent", func(t *testing.T) {
		arts := t.TempDir()
		if err := GenerateCsrf(prepared.Auth{Present: false}, arts); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(arts, "backend")); !os.IsNotExist(err) {
			t.Errorf("expected no output when auth absent")
		}
	})

	t.Run("CookieWritesWithCookieDefault", func(t *testing.T) {
		arts := t.TempDir()
		a := prepared.Auth{CsrfRequired: true, Present: true, Mode: "cookie", Raw: &pmanifest.Auth{}}
		if err := GenerateCsrf(a, arts); err != nil {
			t.Fatalf("error: %v", err)
		}
		body, err := os.ReadFile(filepath.Join(arts, "backend", "internal", "middleware", "csrf.go"))
		if err != nil {
			t.Fatalf("expected csrf.go: %v", err)
		}
		// cookie build regression guard: with BACKEND_AUTH_MODE unset the
		// runtime gate falls back to "cookie", so CSRF stays active.
		if !strings.Contains(string(body), `return "cookie"`) {
			t.Errorf("cookie build must inject \"cookie\" as csrfAuthMode fallback")
		}
	})

	// BUG-116 / Phase-B1 — a manifest=bearer build (CsrfRequired=false) must
	// still write csrf.go, with "bearer" injected as the csrfAuthMode()
	// fallback so the middleware no-ops until BACKEND_AUTH_MODE selects
	// cookie/hybrid at runtime.
	t.Run("BearerWritesRuntimeGated", func(t *testing.T) {
		arts := t.TempDir()
		a := prepared.Auth{CsrfRequired: false, Present: true, Mode: "bearer", Raw: &pmanifest.Auth{Mode: "bearer"}}
		if err := GenerateCsrf(a, arts); err != nil {
			t.Fatalf("error: %v", err)
		}
		body, err := os.ReadFile(filepath.Join(arts, "backend", "internal", "middleware", "csrf.go"))
		if err != nil {
			t.Fatalf("bearer build must still write csrf.go: %v", err)
		}
		for _, must := range []string{`return "bearer"`, "csrfRuntimeActive", "BACKEND_AUTH_MODE"} {
			if !strings.Contains(string(body), must) {
				t.Errorf("bearer csrf.go missing %q", must)
			}
		}
	})

	t.Run("MkdirError", func(t *testing.T) {
		arts := t.TempDir()
		internal := filepath.Join(arts, "backend", "internal")
		if err := os.MkdirAll(filepath.Dir(internal), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.WriteFile(internal, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		a := prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{}}
		if err := GenerateCsrf(a, arts); err == nil {
			t.Errorf("expected mkdir error, got nil")
		}
	})

	t.Run("WriteError", func(t *testing.T) {
		arts := t.TempDir()
		// pre-create csrf.go as a directory so WriteFile fails after mkdir.
		mwDir := filepath.Join(arts, "backend", "internal", "middleware")
		if err := os.MkdirAll(filepath.Join(mwDir, "csrf.go"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		a := prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{}}
		if err := GenerateCsrf(a, arts); err == nil {
			t.Errorf("expected write csrf.go error, got nil")
		}
	})
}
