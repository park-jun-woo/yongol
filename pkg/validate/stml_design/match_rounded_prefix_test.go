//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"testing"
)

func TestMatchRoundedPrefix(t *testing.T) {
	var out pageTokenRefs
	if !matchRoundedPrefix("rounded-card", "rounded-card", "p.stml", &out) {
		t.Fatal("expected rounded-card to match")
	}
	if len(out.Rounded) != 1 || out.Rounded[0].Name != "card" {
		t.Errorf("rounded = %+v", out.Rounded)
	}
	var out2 pageTokenRefs
	if matchRoundedPrefix("rounded-full", "rounded-full", "p.stml", &out2) {
		t.Error("rounded-full is builtin, should be skipped")
	}
	if matchRoundedPrefix("flex", "flex", "p.stml", &out2) {
		t.Error("flex not a rounded prefix")
	}
}
