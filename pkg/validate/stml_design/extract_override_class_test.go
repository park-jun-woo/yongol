//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"testing"
)

func TestExtractOverrideClass(t *testing.T) {
	if got := extractOverrideClass(`@override class="card primary"`); got != "card primary" {
		t.Errorf("got %q, want 'card primary'", got)
	}
	if got := extractOverrideClass(`@override class='solo'`); got != "solo" {
		t.Errorf("single quote: %q", got)
	}
	if got := extractOverrideClass(`@override`); got != "" {
		t.Errorf("structure-only override should yield empty, got %q", got)
	}
}
