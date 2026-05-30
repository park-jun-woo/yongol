//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCopyFuncSpecs — specs/func Go 파일 복사 + skip/non-go/walk-error 분기 검증

package gogin

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFuncSpecs(t *testing.T) {
	t.Run("MissingFuncDirSkips", func(t *testing.T) {
		specs := t.TempDir()
		arts := t.TempDir()
		if err := copyFuncSpecs(specs, arts); err != nil {
			t.Errorf("expected nil for missing func dir, got: %v", err)
		}
		if _, err := os.Stat(filepath.Join(arts, "backend")); !os.IsNotExist(err) {
			t.Errorf("expected no output, stat err: %v", err)
		}
	})

	t.Run("CopiesGoFilesOnly", func(t *testing.T) {
		specs := t.TempDir()
		arts := t.TempDir()
		pkgDir := filepath.Join(specs, "func", "billing")
		if err := os.MkdirAll(pkgDir, 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		mustWrite(t, filepath.Join(pkgDir, "deduct.go"), "package billing")
		mustWrite(t, filepath.Join(pkgDir, "notes.txt"), "skip me")

		if err := copyFuncSpecs(specs, arts); err != nil {
			t.Fatalf("copyFuncSpecs error: %v", err)
		}
		dst := filepath.Join(arts, "backend", "internal", "billing")
		if _, err := os.Stat(filepath.Join(dst, "deduct.go")); err != nil {
			t.Errorf("expected deduct.go copied: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dst, "notes.txt")); !os.IsNotExist(err) {
			t.Errorf("non-.go file should not be copied, stat err: %v", err)
		}
	})

	t.Run("WalkError", func(t *testing.T) {
		specs := t.TempDir()
		arts := t.TempDir()
		funcDir := filepath.Join(specs, "func")
		// A 0-perm subdir with contents makes Walk's ReadDir of it fail,
		// surfacing an error to the walk callback (err != nil branch).
		bad := filepath.Join(funcDir, "noaccess")
		if err := os.MkdirAll(filepath.Join(bad, "inner"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.Chmod(bad, 0o000); err != nil {
			t.Fatalf("setup: %v", err)
		}
		defer os.Chmod(bad, 0o755)
		err := copyFuncSpecs(specs, arts)
		if err == nil {
			t.Skip("walk did not surface a permission error (likely running as root)")
		}
	})
}
