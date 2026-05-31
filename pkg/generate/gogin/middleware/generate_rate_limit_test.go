//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMiddlewareExtras — csrf/prometheus/rate-limit/request-id/writeFiles 유틸 검증
package middleware

import (
	"os"
	"path/filepath"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateRateLimit(t *testing.T) {
	t.Run("SkipsNilManifest", func(t *testing.T) {
		if err := GenerateRateLimit(&yongol.Fullstack{}, t.TempDir()); err != nil {
			t.Errorf("error: %v", err)
		}
	})
	t.Run("SkipsEmptyModule", func(t *testing.T) {
		fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
		if err := GenerateRateLimit(fs, t.TempDir()); err != nil {
			t.Errorf("error: %v", err)
		}
	})
	t.Run("WritesFiles", func(t *testing.T) {
		arts := t.TempDir()
		fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}}}
		if err := GenerateRateLimit(fs, arts); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(arts, "backend", "internal", "middleware", "fixed_rate_limit.go")); err != nil {
			t.Errorf("expected fixed_rate_limit.go: %v", err)
		}
	})
}
