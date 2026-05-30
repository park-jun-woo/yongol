//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestFuncSpecRelPaths — specs/func 하위 디렉토리 → internal/<pkg> 경로 + skip/error 검증

package gogin

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestFuncSpecRelPaths(t *testing.T) {
	t.Run("MissingDirReturnsNil", func(t *testing.T) {
		got, err := funcSpecRelPaths(t.TempDir())
		if err != nil {
			t.Fatalf("expected nil error, got: %v", err)
		}
		if got != nil {
			t.Errorf("expected nil slice, got: %v", got)
		}
	})

	t.Run("ListsSubdirsOnly", func(t *testing.T) {
		specs := t.TempDir()
		funcDir := filepath.Join(specs, "func")
		if err := os.MkdirAll(filepath.Join(funcDir, "auth"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.MkdirAll(filepath.Join(funcDir, "billing"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		mustWrite(t, filepath.Join(funcDir, "readme.txt"), "x") // file -> skipped

		got, err := funcSpecRelPaths(specs)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		sort.Strings(got)
		want := []string{"internal/auth", "internal/billing"}
		if len(got) != len(want) {
			t.Fatalf("expected %v, got %v", want, got)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("got[%d]=%q, want %q", i, got[i], want[i])
			}
		}
	})

	t.Run("ReadDirError", func(t *testing.T) {
		specs := t.TempDir()
		// specs/func is a regular file -> ReadDir returns a non-IsNotExist error.
		if err := os.WriteFile(filepath.Join(specs, "func"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if _, err := funcSpecRelPaths(specs); err == nil {
			t.Errorf("expected ReadDir error, got nil")
		}
	})
}
