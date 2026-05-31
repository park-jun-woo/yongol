//ff:func feature=statemachine type=test control=sequence
//ff:what extractMermaidBlock / extractMermaidBlockWithLine / findStateLine / checkCaseConflicts
package statemachine

import (
	"testing"
)

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
