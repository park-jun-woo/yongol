//ff:func feature=statemachine type=test control=iteration dimension=1
//ff:what extractMermaidBlock / extractMermaidBlockWithLine / findStateLine / checkCaseConflicts

package statemachine

import "testing"

func TestExtractMermaidBlock(t *testing.T) {
	content := "# Title\n\n```mermaid\nstateDiagram\n  A --> B\n```\nrest"
	got := extractMermaidBlock(content)
	if got != "\nstateDiagram\n  A --> B\n" {
		t.Errorf("extractMermaidBlock = %q", got)
	}
	// no block
	if got := extractMermaidBlock("no block here"); got != "" {
		t.Errorf("no block = %q, want empty", got)
	}
	// unterminated block
	if got := extractMermaidBlock("```mermaid\nunterminated"); got != "" {
		t.Errorf("unterminated = %q, want empty", got)
	}
}

func TestExtractMermaidBlockWithLine(t *testing.T) {
	content := "line1\nline2\n```mermaid\nA --> B\n```"
	body, line := extractMermaidBlockWithLine(content)
	if line != 3 {
		t.Errorf("line = %d, want 3", line)
	}
	if body != "\nA --> B\n" {
		t.Errorf("body = %q", body)
	}
	// no block
	if b, l := extractMermaidBlockWithLine("none"); b != "" || l != 0 {
		t.Errorf("no block = (%q,%d)", b, l)
	}
	// unterminated
	if b, l := extractMermaidBlockWithLine("```mermaid\nX"); b != "" || l != 0 {
		t.Errorf("unterminated = (%q,%d)", b, l)
	}
}

func TestFindStateLine(t *testing.T) {
	lines := []string{"stateDiagram", "  Pending --> Active", "  Active --> Done"}
	// "Active" first appears on lines[1] (index 1); mermaidStartLine=5 → 5+1+1=7
	if got := findStateLine(lines, "Active", 5); got != 7 {
		t.Errorf("findStateLine = %d, want 7", got)
	}
	// state not found returns mermaidStartLine
	if got := findStateLine(lines, "Missing", 5); got != 5 {
		t.Errorf("missing state = %d, want 5", got)
	}
}

func TestCheckCaseConflicts(t *testing.T) {
	lines := []string{"Active --> active"}
	stateSet := map[string]bool{"Active": true, "active": true}
	diags := checkCaseConflicts("courseSM", "x.md", stateSet, lines, 10)
	if len(diags) != 1 {
		t.Fatalf("expected 1 conflict diag, got %d (%+v)", len(diags), diags)
	}
	if diags[0].File != "x.md" {
		t.Errorf("diag file = %q", diags[0].File)
	}

	// no conflict when states differ beyond case
	noConflict := map[string]bool{"Active": true, "Pending": true}
	if d := checkCaseConflicts("sm", "x.md", noConflict, lines, 0); len(d) != 0 {
		t.Errorf("expected no conflicts, got %+v", d)
	}
}
