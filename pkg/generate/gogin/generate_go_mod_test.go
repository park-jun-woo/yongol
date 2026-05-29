//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRunGoModTidySuccess — 최소 유효 go.mod 에서 runGoModTidy 성공 확인
package gogin

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// TestRunGoModTidySuccess verifies that runGoModTidy succeeds on a minimal
// valid go.mod (no require directives — nothing to resolve).
func TestRunGoModTidySuccess(t *testing.T) {
	if _, err := exec.LookPath("go"); err != nil {
		t.Skip("go toolchain not available; skipping")
	}
	dir := t.TempDir()
	modContent := "module example.com/phase010/success\n\ngo 1.22\n"
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(modContent), 0o644); err != nil {
		t.Fatalf("write go.mod: %v", err)
	}
	if err := runGoModTidy(dir); err != nil {
		t.Fatalf("runGoModTidy unexpected error: %v", err)
	}
}
