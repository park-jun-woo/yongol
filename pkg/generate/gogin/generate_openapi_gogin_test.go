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
		// backend is a regular file -> backend/internal/api MkdirAll fails.
		if err := os.WriteFile(filepath.Join(arts, "backend"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		err := generateOpenAPIGoGin(t.TempDir(), arts)
		if err == nil || !strings.Contains(err.Error(), "mkdir") {
			t.Errorf("expected mkdir error, got: %v", err)
		}
	})

	t.Run("ExecFailsOnMissingSpec", func(t *testing.T) {
		// specs/api/openapi.yaml does not exist -> oapi-codegen exits non-zero
		// (or the binary is absent); either way exec returns an error.
		err := generateOpenAPIGoGin(t.TempDir(), t.TempDir())
		if err == nil {
			t.Errorf("expected exec error for missing openapi.yaml, got nil")
		}
	})
}
