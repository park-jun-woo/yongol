//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateOpenAPIGoGin — mkdir 에러 + oapi-codegen 실행 실패 경로 검증

package gogin

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateOpenAPIGoGin(t *testing.T) {
	t.Run("MkdirError", func(t *testing.T) {
		arts := t.TempDir()
		// parent is a regular file -> outDir MkdirAll fails.
		parent := filepath.Join(arts, "backend")
		if err := os.WriteFile(parent, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		outDir := filepath.Join(parent, "internal", "api")
		err := generateOpenAPIGoGin(filepath.Join(t.TempDir(), "openapi.yaml"), outDir, "api")
		if err == nil || !strings.Contains(err.Error(), "mkdir") {
			t.Errorf("expected mkdir error, got: %v", err)
		}
	})

	t.Run("ExecFailsOnMissingSpec", func(t *testing.T) {
		// the spec path does not exist -> oapi-codegen exits non-zero
		// (or the binary is absent); either way exec returns an error.
		oapiPath := filepath.Join(t.TempDir(), "api", "openapi.yaml")
		err := generateOpenAPIGoGin(oapiPath, t.TempDir(), "api")
		if err == nil {
			t.Errorf("expected exec error for missing openapi.yaml, got nil")
		}
	})
}
