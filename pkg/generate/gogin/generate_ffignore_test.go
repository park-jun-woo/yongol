//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateFFIgnore — backend/.ffignore 생성 success + error 경로 검증

package gogin

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateFFIgnore(t *testing.T) {
	t.Run("WritesFFIgnore", func(t *testing.T) {
		arts := t.TempDir()
		if err := os.MkdirAll(filepath.Join(arts, "backend"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := generateFFIgnore(arts); err != nil {
			t.Fatalf("generateFFIgnore error: %v", err)
		}
		raw, err := os.ReadFile(filepath.Join(arts, "backend", ".ffignore"))
		if err != nil {
			t.Fatalf("read .ffignore: %v", err)
		}
		got := string(raw)
		if !strings.Contains(got, "internal/db/db.go") {
			t.Errorf("expected sqlc db.go entry, got:\n%s", got)
		}
		if !strings.Contains(got, "register_handlers_with_options.gen.go") {
			t.Errorf("expected oapi-codegen entry, got:\n%s", got)
		}
	})

	t.Run("WriteError", func(t *testing.T) {
		// backend dir does not exist -> WriteFile fails (no MkdirAll in func).
		arts := t.TempDir()
		if err := generateFFIgnore(arts); err == nil {
			t.Errorf("expected write error when backend dir absent, got nil")
		}
	})
}
