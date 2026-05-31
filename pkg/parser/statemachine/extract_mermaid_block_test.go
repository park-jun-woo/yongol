//ff:func feature=statemachine type=test control=sequence
//ff:what extractMermaidBlock / extractMermaidBlockWithLine / findStateLine / checkCaseConflicts
package statemachine

import (
	"testing"
)

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
