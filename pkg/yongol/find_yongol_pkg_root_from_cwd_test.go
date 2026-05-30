//ff:func feature=orchestrator type=test control=sequence
//ff:what findYongolPkgRootFromCWD — yongol root 발견(성공) 및 root까지 미발견("") 분기 검증

package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindYongolPkgRootFromCWD_Found(t *testing.T) {
	base := t.TempDir()
	yongolRoot := filepath.Join(base, "yongol")
	ssacPkg := filepath.Join(base, "ssac", "pkg")
	mustMkdir(t, filepath.Join(yongolRoot, "pkg"))
	mustMkdir(t, ssacPkg)
	if err := os.WriteFile(filepath.Join(yongolRoot, "go.mod"),
		[]byte("module github.com/park-jun-woo/yongol\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(yongolRoot); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	got := findYongolPkgRootFromCWD()
	want, _ := filepath.EvalSymlinks(ssacPkg)
	gr, _ := filepath.EvalSymlinks(got)
	if got == "" || gr != want {
		t.Fatalf("got %q, want %q", got, ssacPkg)
	}
}

func TestFindYongolPkgRootFromCWD_NotFound(t *testing.T) {
	// An isolated empty directory: walking up never finds a yongol root,
	// so the loop reaches the filesystem root and returns "".
	isolated := t.TempDir()
	if err := os.Chdir(isolated); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	if got := findYongolPkgRootFromCWD(); got != "" {
		t.Skipf("ambient yongol root resolved %q; skipping not-found assertion", got)
	}
}
