//ff:func feature=statemachine type=test control=sequence
//ff:what extractMermaidBlock / extractMermaidBlockWithLine / findStateLine / checkCaseConflicts
package statemachine

import (
	"testing"
)

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
