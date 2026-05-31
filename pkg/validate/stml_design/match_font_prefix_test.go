//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"testing"
)

func TestMatchFontPrefix(t *testing.T) {
	var out pageTokenRefs
	matchFontPrefix("font-display", "font-display", "p.stml", &out)
	if len(out.Fonts) != 1 || out.Fonts[0].Name != "display" {
		t.Errorf("fonts = %+v", out.Fonts)
	}
	// Skippable builtin → no record.
	var out2 pageTokenRefs
	matchFontPrefix("font-sans", "font-sans", "p.stml", &out2)
	if len(out2.Fonts) != 0 {
		t.Errorf("font-sans is builtin, should be skipped: %+v", out2.Fonts)
	}
	// Non-font prefix → no record.
	matchFontPrefix("text-lg", "text-lg", "p.stml", &out2)
	if len(out2.Fonts) != 0 {
		t.Error("text-lg should not record a font")
	}
}
