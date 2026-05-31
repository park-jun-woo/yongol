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

func TestGeneratePrometheus(t *testing.T) {
	t.Run("SkipsNilManifest", func(t *testing.T) {
		if err := GeneratePrometheus(&yongol.Fullstack{}, t.TempDir()); err != nil {
			t.Errorf("error: %v", err)
		}
	})
	t.Run("SkipsEmptyModule", func(t *testing.T) {
		fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
		if err := GeneratePrometheus(fs, t.TempDir()); err != nil {
			t.Errorf("error: %v", err)
		}
	})
	t.Run("WritesFiles", func(t *testing.T) {
		arts := t.TempDir()
		fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}}}
		if err := GeneratePrometheus(fs, arts); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(arts, "backend", "internal", "middleware", "prometheus_middleware.go")); err != nil {
			t.Errorf("expected prometheus_middleware.go: %v", err)
		}
	})
}
