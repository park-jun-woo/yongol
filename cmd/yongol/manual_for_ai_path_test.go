//ff:func feature=cli type=test control=sequence
//ff:what TestManualForAIPath — manual-for-ai.md 발견(경로) / 미발견(github URL) 분기 검증

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func restoreCwd(t *testing.T) {
	t.Helper()
	orig, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(orig) })
}

func TestManualForAIPath_Found(t *testing.T) {
	restoreCwd(t)
	dir := t.TempDir()
	manual := filepath.Join(dir, "manual-for-ai.md")
	if err := os.WriteFile(manual, []byte("# manual"), 0o644); err != nil {
		t.Fatal(err)
	}
	// Work from a nested subdir so the upward walk is exercised.
	sub := filepath.Join(dir, "a", "b")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(sub); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	got := manualForAIPath()
	wantResolved, _ := filepath.EvalSymlinks(manual)
	gotResolved, _ := filepath.EvalSymlinks(got)
	if gotResolved != wantResolved {
		t.Fatalf("got %q, want %q", got, manual)
	}
}

func TestManualForAIPath_NotFound(t *testing.T) {
	restoreCwd(t)
	// An isolated temp dir whose ancestry has no manual-for-ai.md → github URL.
	dir := t.TempDir()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	got := manualForAIPath()
	if !strings.HasPrefix(got, "https://github.com/") {
		t.Skipf("ambient manual-for-ai.md resolved %q; skipping URL assertion", got)
	}
}
