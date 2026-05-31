//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"testing"
)

func TestMatchColorPrefix(t *testing.T) {
	var out pageTokenRefs
	if !matchColorPrefix("bg-primary", "bg-primary", "p.stml", &out) {
		t.Fatal("expected bg-primary to match color")
	}
	if len(out.Colors) != 1 || out.Colors[0].Name != "primary" {
		t.Errorf("colors = %+v", out.Colors)
	}
	// Skippable name (numeric) → no record, no match.
	var out2 pageTokenRefs
	if matchColorPrefix("bg-500", "bg-500", "p.stml", &out2) {
		// matched? actually "500" -> isPureNumeric true -> continue -> returns false
		t.Error("numeric color value should be skippable")
	}
	// Non-color prefix → false.
	if matchColorPrefix("flex", "flex", "p.stml", &out2) {
		t.Error("flex should not match a color prefix")
	}
}
