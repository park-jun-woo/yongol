//ff:func feature=statemachine type=test control=iteration dimension=1 topic=states
//ff:what ParseDir — 정상/비정상 mermaid 파일 혼재 시 부분 성공 (정상 diagram + 에러 diag)

package statemachine

import (
	"os"
	"path/filepath"
	"testing"
)

// TestParseDirPartialSuccessMixed verifies that a directory with one valid
// and one broken mermaid .md file returns both a parsed diagram and a diag.
func TestParseDirPartialSuccessMixed(t *testing.T) {
	dir := t.TempDir()

	okPath := filepath.Join(dir, "gig.md")
	if err := os.WriteFile(okPath, []byte(validMermaidBody), 0o644); err != nil {
		t.Fatalf("write ok file: %v", err)
	}

	badPath := filepath.Join(dir, "broken.md")
	// No mermaid fenced block → Parse emits a LevelError diagnostic.
	if err := os.WriteFile(badPath, []byte("# No mermaid block here\n"), 0o644); err != nil {
		t.Fatalf("write bad file: %v", err)
	}

	diagrams, diags := ParseDir(dir)

	if len(diagrams) < 1 {
		t.Fatalf("expected at least 1 diagram, got %d (diags: %v)", len(diagrams), diags)
	}
	if len(diags) < 1 {
		t.Fatalf("expected at least 1 diag for broken file, got 0")
	}

	var found bool
	for _, d := range diagrams {
		if d.ID != "gig" {
			continue
		}
		found = true
		if len(d.Transitions) == 0 {
			t.Errorf("gig diagram has no transitions")
		}
	}
	if !found {
		t.Errorf("successful diagram 'gig' missing from diagrams slice")
	}
}
