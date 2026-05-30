//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRunGoModTidySuccess — 최소 유효 go.mod 에서 runGoModTidy 성공 확인
package gogin

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestGenerateGoMod_MkdirError verifies that generateGoMod surfaces a wrapped
// error when the backend output directory cannot be created (its parent path
// already exists as a regular file). This path is deterministic and needs no
// go toolchain / network access.
func TestGenerateGoMod_MkdirError(t *testing.T) {
	dir := t.TempDir()
	// artifactsDir is a regular file -> MkdirAll(artifactsDir/backend) fails.
	blocker := filepath.Join(dir, "blocker")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}},
	}
	err := generateGoMod(fs, "example.com/app", blocker)
	if err == nil {
		t.Fatalf("expected mkdir error when backend parent is a file, got nil")
	}
}

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
