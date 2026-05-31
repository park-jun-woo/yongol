//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"testing"
)

func TestClassifySingleToken(t *testing.T) {
	// Responsive prefix stripped, then color matched.
	var out pageTokenRefs
	classifySingleToken("sm:bg-primary", "sm:bg-primary", "p.stml", &out)
	if len(out.Colors) != 1 || out.Colors[0].Name != "primary" {
		t.Errorf("responsive color = %+v", out.Colors)
	}
	// Negative spacing prefix.
	var out2 pageTokenRefs
	classifySingleToken("-mt-gutter", "-mt-gutter", "p.stml", &out2)
	if len(out2.Spacing) != 1 || out2.Spacing[0].Name != "gutter" {
		t.Errorf("negative spacing = %+v", out2.Spacing)
	}
	// Font fallback.
	var out3 pageTokenRefs
	classifySingleToken("font-brand", "font-brand", "p.stml", &out3)
	if len(out3.Fonts) != 1 {
		t.Errorf("font = %+v", out3.Fonts)
	}
}
