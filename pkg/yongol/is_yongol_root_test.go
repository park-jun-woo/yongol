//ff:func feature=orchestrator type=test control=sequence
//ff:what TestIsDir/dirPresence/isYongolRoot — 디렉토리 판별·presence 매핑·루트 판별 검증
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

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
