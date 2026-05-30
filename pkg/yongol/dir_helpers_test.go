//ff:func feature=orchestrator type=test control=sequence
//ff:what TestIsDir/dirPresence/isYongolRoot — 디렉토리 판별·presence 매핑·루트 판별 검증

package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsDir(t *testing.T) {
	dir := t.TempDir()
	if !isDir(dir) {
		t.Error("expected isDir(tempdir) = true")
	}

	file := filepath.Join(dir, "f.txt")
	if err := os.WriteFile(file, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if isDir(file) {
		t.Error("expected isDir(file) = false")
	}
	if isDir(filepath.Join(dir, "missing")) {
		t.Error("expected isDir(missing) = false")
	}
}

func TestDirPresence(t *testing.T) {
	dir := t.TempDir()

	if got := dirPresence(dir, 3); got != SSOTPopulated {
		t.Errorf("dirPresence(existing, 3) = %v, want SSOTPopulated", got)
	}
	if got := dirPresence(dir, 0); got != SSOTDeclared {
		t.Errorf("dirPresence(existing, 0) = %v, want SSOTDeclared", got)
	}
	missing := filepath.Join(dir, "nope")
	if got := dirPresence(missing, 0); got != SSOTAbsent {
		t.Errorf("dirPresence(missing, 0) = %v, want SSOTAbsent", got)
	}
	// File count > 0 wins even when the directory does not exist.
	if got := dirPresence(missing, 1); got != SSOTPopulated {
		t.Errorf("dirPresence(missing, 1) = %v, want SSOTPopulated", got)
	}
}

func TestIsYongolRoot(t *testing.T) {
	// Valid root: go.mod with module path + pkg/ dir.
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "go.mod"),
		[]byte("module github.com/park-jun-woo/yongol\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(root, "pkg"), 0o755); err != nil {
		t.Fatal(err)
	}
	if !isYongolRoot(root) {
		t.Error("expected isYongolRoot(valid) = true")
	}

	// Missing go.mod.
	if isYongolRoot(t.TempDir()) {
		t.Error("expected isYongolRoot(no go.mod) = false")
	}

	// go.mod present but wrong module path.
	wrong := t.TempDir()
	if err := os.WriteFile(filepath.Join(wrong, "go.mod"),
		[]byte("module example.com/other\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(wrong, "pkg"), 0o755); err != nil {
		t.Fatal(err)
	}
	if isYongolRoot(wrong) {
		t.Error("expected isYongolRoot(wrong module) = false")
	}

	// Correct module path but no pkg/ dir.
	noPkg := t.TempDir()
	if err := os.WriteFile(filepath.Join(noPkg, "go.mod"),
		[]byte("module github.com/park-jun-woo/yongol\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if isYongolRoot(noPkg) {
		t.Error("expected isYongolRoot(no pkg dir) = false")
	}
}
