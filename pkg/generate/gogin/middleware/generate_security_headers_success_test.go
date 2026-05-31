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

func TestGenerateSecurityHeaders_Success(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}}}
	arts := t.TempDir()
	if err := GenerateSecurityHeaders(fs, arts); err != nil {
		t.Fatalf("GenerateSecurityHeaders: %v", err)
	}
	mwDir := filepath.Join(arts, "backend", "internal", "middleware")
	if _, err := os.Stat(filepath.Join(mwDir, "security_headers_config.go")); err != nil {
		t.Errorf("expected security_headers_config.go: %v", err)
	}
}
