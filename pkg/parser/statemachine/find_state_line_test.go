//ff:func feature=statemachine type=test control=sequence
//ff:what extractMermaidBlock / extractMermaidBlockWithLine / findStateLine / checkCaseConflicts
package statemachine

import (
	"testing"
)

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
