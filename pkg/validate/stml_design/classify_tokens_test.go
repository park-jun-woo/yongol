//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"testing"
)

func TestClassifyTokens(t *testing.T) {
	var out pageTokenRefs
	classifyTokens("bg-primary rounded-card p-section", "p.stml", &out)
	if len(out.Colors) != 1 || len(out.Rounded) != 1 || len(out.Spacing) != 1 {
		t.Errorf("classifyTokens result: %+v", out)
	}
	// Empty class is a no-op.
	var out2 pageTokenRefs
	classifyTokens("", "p.stml", &out2)
	if len(out2.Colors)+len(out2.Spacing)+len(out2.Rounded)+len(out2.Fonts) != 0 {
		t.Error("empty class should produce nothing")
	}
}
