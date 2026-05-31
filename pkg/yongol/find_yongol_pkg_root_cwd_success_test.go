//ff:func feature=orchestrator type=test control=sequence
//ff:what findYongolPkgRoot — env 미설정 시 CWD fallback이 sibling ssac/pkg를 반환
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindYongolPkgRootCWDSuccess(t *testing.T) {
	base := t.TempDir()
	yongolRoot := filepath.Join(base, "yongol")
	ssacPkg := filepath.Join(base, "ssac", "pkg")

	mustMkdir(t, filepath.Join(yongolRoot, "pkg"))
	mustMkdir(t, ssacPkg)
	if err := os.WriteFile(filepath.Join(yongolRoot, "go.mod"),
		[]byte("module github.com/park-jun-woo/yongol\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	// Ensure env override is absent so the CWD branch is exercised.
	t.Setenv("YONGOL_SSAC_PKG", filepath.Join(base, "nope"))
	if err := os.Chdir(yongolRoot); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	got := findYongolPkgRoot()
	// EvalSymlinks because t.TempDir may live under a symlinked /tmp.
	wantResolved, _ := filepath.EvalSymlinks(ssacPkg)
	gotResolved, _ := filepath.EvalSymlinks(got)
	if got == "" || gotResolved != wantResolved {
		t.Fatalf("findYongolPkgRoot = %q; want %q", got, ssacPkg)
	}
}
