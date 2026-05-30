//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWriteEnvHelperFiles — cmd/<name>.go 헬퍼 emit + skip/error 경로 검증

package boot

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteEnvHelperFiles(t *testing.T) {
	funcs := []string{
		"func envInt(key string, def int) int {\n\treturn def\n}\n",
		"   ",              // empty after trim -> skipped
		"var notAFunc = 1", // no "func " prefix -> parseFuncName "" -> skipped
	}

	t.Run("WritesNamedHelperOnly", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeEnvHelperFiles(dir, []string{`"os"`}, funcs); err != nil {
			t.Fatalf("writeEnvHelperFiles error: %v", err)
		}
		cmdDir := filepath.Join(dir, "backend", "cmd")
		entries, err := os.ReadDir(cmdDir)
		if err != nil {
			t.Fatalf("read cmd dir: %v", err)
		}
		// Only the valid func should have produced a file.
		if len(entries) != 1 {
			t.Errorf("expected exactly 1 helper file, got %d: %v", len(entries), entries)
		}
		if _, err := os.Stat(filepath.Join(cmdDir, "env_int.go")); err != nil {
			t.Errorf("expected env_int.go: %v", err)
		}
	})

	t.Run("MkdirError", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "backend"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeEnvHelperFiles(dir, nil, funcs); err == nil {
			t.Errorf("expected MkdirAll error, got nil")
		}
	})

	t.Run("WriteFileError", func(t *testing.T) {
		dir := t.TempDir()
		cmdDir := filepath.Join(dir, "backend", "cmd")
		if err := os.MkdirAll(filepath.Join(cmdDir, "env_int.go"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeEnvHelperFiles(dir, nil, funcs); err == nil {
			t.Errorf("expected WriteFile error, got nil")
		}
	})
}
