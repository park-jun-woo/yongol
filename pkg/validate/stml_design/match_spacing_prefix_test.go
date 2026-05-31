//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"testing"
)

func TestMatchSpacingPrefix(t *testing.T) {
	var out pageTokenRefs
	if !matchSpacingPrefix("p-section", "p-section", "p.stml", &out) {
		t.Fatal("expected p-section to match spacing")
	}
	if len(out.Spacing) != 1 || out.Spacing[0].Name != "section" {
		t.Errorf("spacing = %+v", out.Spacing)
	}
	var out2 pageTokenRefs
	if matchSpacingPrefix("p-4", "p-4", "p.stml", &out2) {
		t.Error("p-4 numeric should be skippable")
	}
	if matchSpacingPrefix("flex", "flex", "p.stml", &out2) {
		t.Error("flex not a spacing prefix")
	}
}
