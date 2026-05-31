//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"testing"
)

func TestIsOverridden(t *testing.T) {
	ovr := overrideSet{
		"page.stml": {"card": true},
	}
	if !isOverridden(ovr, "page.stml", "card") {
		t.Error("expected card overridden")
	}
	if isOverridden(ovr, "page.stml", "button") {
		t.Error("button not overridden")
	}
	if isOverridden(ovr, "other.stml", "card") {
		t.Error("unknown file not overridden")
	}
	if isOverridden(ovr, "page.stml", "") {
		t.Error("empty class never overridden")
	}
}
