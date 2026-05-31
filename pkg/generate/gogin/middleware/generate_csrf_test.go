//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMiddlewareExtras — csrf/prometheus/rate-limit/request-id/writeFiles 유틸 검증
package middleware

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestGenerateCsrf(t *testing.T) {
	t.Run("SkipsWhenInactive", func(t *testing.T) {
		arts := t.TempDir()
		if err := GenerateCsrf(prepared.Auth{CsrfRequired: false}, arts); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(arts, "backend")); !os.IsNotExist(err) {
			t.Errorf("expected no output when csrf inactive")
		}
	})

	t.Run("WritesWhenActive", func(t *testing.T) {
		arts := t.TempDir()
		a := prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{}}
		if err := GenerateCsrf(a, arts); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(arts, "backend", "internal", "middleware", "csrf.go")); err != nil {
			t.Errorf("expected csrf.go: %v", err)
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
