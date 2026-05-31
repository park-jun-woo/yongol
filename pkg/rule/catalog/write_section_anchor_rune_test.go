//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what TestCatalogHelpers — unit tests for the pure rule catalog helper functions
package catalog

import (
	"strings"
	"testing"
)

func TestWriteSectionAnchorRune(t *testing.T) {
	// Alphanumeric is written, resets prevDash.
	var b strings.Builder
	if dash := writeSectionAnchorRune(&b, 'a', false); dash {
		t.Error("alnum should return prevDash=false")
	}
	if b.String() != "a" {
		t.Errorf("buffer = %q", b.String())
	}
	// Space after content writes a dash.
	if dash := writeSectionAnchorRune(&b, ' ', false); !dash {
		t.Error("space after content should write dash")
	}
	if b.String() != "a-" {
		t.Errorf("buffer = %q, want a-", b.String())
	}
	// Leading separator (empty buffer) is suppressed.
	var b2 strings.Builder
	if dash := writeSectionAnchorRune(&b2, '-', false); dash {
		t.Error("leading separator should not write dash")
	}
	if b2.Len() != 0 {
		t.Errorf("leading separator wrote %q", b2.String())
	}
	// Unknown punctuation is dropped, prevDash preserved.
	var b3 strings.Builder
	b3.WriteString("x")
	if dash := writeSectionAnchorRune(&b3, '!', false); dash {
		t.Error("punctuation should preserve prevDash=false")
	}
	if b3.String() != "x" {
		t.Errorf("punctuation altered buffer: %q", b3.String())
	}
}
