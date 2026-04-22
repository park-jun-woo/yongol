//ff:func feature=statemachine type=test control=iteration dimension=1 topic=states
//ff:what ParseDir partial success 검증 — mixed / all-ok / missing-dir

package statemachine

import (
	"os"
	"path/filepath"
	"testing"
)

const validMermaidBody = "# Gig\n\n```mermaid\nstateDiagram-v2\n    [*] --> draft\n    draft --> open: PublishGig\n    open --> closed: CloseGig\n```\n"

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
		if d.ID == "gig" {
			found = true
			if len(d.Transitions) == 0 {
				t.Errorf("gig diagram has no transitions")
			}
		}
	}
	if !found {
		t.Errorf("successful diagram 'gig' missing from diagrams slice")
	}
}

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

// TestParseDirPartialSuccessMissingDir verifies SILENT-OK behaviour for a
// missing states/ directory (optional SSOT): diags nil, diagrams nil.
func TestParseDirPartialSuccessMissingDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "does-not-exist")

	diagrams, diags := ParseDir(dir)

	if len(diags) != 0 {
		t.Fatalf("expected 0 diags for missing dir (SILENT-OK), got %d: %v", len(diags), diags)
	}
	if len(diagrams) != 0 {
		t.Fatalf("expected 0 diagrams for missing dir, got %d", len(diagrams))
	}
}
