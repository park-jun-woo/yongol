//ff:func feature=agent type=test control=sequence
//ff:what TestFindSSaCFile — service/*/{op}.ssac 글로브 탐색 및 부재 시 false 검증
package agent

import (
	"os"
	"path/filepath"
	"testing"
)

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
