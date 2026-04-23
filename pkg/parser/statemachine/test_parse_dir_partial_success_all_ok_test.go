//ff:func feature=statemachine type=test control=iteration dimension=1 topic=states
//ff:what ParseDir — 모든 파일 정상이면 진단 0개, diagram 전부 수집

package statemachine

import (
	"os"
	"path/filepath"
	"testing"
)

// TestParseDirPartialSuccessAllOK is a regression test: no diags when every
// .md file parses cleanly.
func TestParseDirPartialSuccessAllOK(t *testing.T) {
	dir := t.TempDir()

	for _, name := range []string{"a.md", "b.md"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(validMermaidBody), 0o644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}

	diagrams, diags := ParseDir(dir)

	if len(diags) != 0 {
		t.Fatalf("expected 0 diags for all-ok dir, got %d: %v", len(diags), diags)
	}
	if len(diagrams) != 2 {
		t.Fatalf("expected 2 diagrams, got %d", len(diagrams))
	}
}
