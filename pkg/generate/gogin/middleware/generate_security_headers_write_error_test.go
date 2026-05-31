//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what TestGenerateSecurityHeaders — nil/empty-module skip + 성공 + writeFiles 에러 분기
package middleware

import (
	"os"
	"path/filepath"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateSecurityHeaders_WriteError(t *testing.T) {
	arts := t.TempDir()
	// Place a regular file where the middleware directory must be created.
	blocker := filepath.Join(arts, "backend", "internal")
	if err := os.MkdirAll(blocker, 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := os.WriteFile(filepath.Join(blocker, "middleware"), []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}}}
	if err := GenerateSecurityHeaders(fs, arts); err == nil {
		t.Errorf("expected write error, got nil")
	}
}
