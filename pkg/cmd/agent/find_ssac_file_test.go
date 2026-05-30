//ff:func feature=agent type=test control=sequence
//ff:what TestFindSSaCFile — service/*/{op}.ssac 글로브 탐색 및 부재 시 false 검증

package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindSSaCFile(t *testing.T) {
	dir := t.TempDir()
	svc := filepath.Join(dir, "service", "auth")
	if err := os.MkdirAll(svc, 0o755); err != nil {
		t.Fatal(err)
	}
	body := "func Login() {}"
	if err := os.WriteFile(filepath.Join(svc, "Login.ssac"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}

	content, ok := findSSaCFile(dir, "Login")
	if !ok {
		t.Fatal("expected to find Login.ssac")
	}
	if content != body {
		t.Errorf("content = %q, want %q", content, body)
	}

	if _, ok := findSSaCFile(dir, "Missing"); ok {
		t.Error("expected false for missing operationId")
	}
}

func TestFindSSaCFileReadError(t *testing.T) {
	// The glob matches a *directory* named "Dir.ssac"; os.ReadFile on a directory
	// fails, exercising the read-error branch.
	dir := t.TempDir()
	svc := filepath.Join(dir, "service", "auth")
	if err := os.MkdirAll(filepath.Join(svc, "Dir.ssac"), 0o755); err != nil {
		t.Fatal(err)
	}
	if content, ok := findSSaCFile(dir, "Dir"); ok || content != "" {
		t.Errorf("expected (\"\", false) when match is unreadable, got (%q, %v)", content, ok)
	}
}
