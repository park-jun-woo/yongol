//ff:func feature=stml-parse type=test control=sequence
//ff:what isImplicitTag/isDataElement/getAttr/hasAttr/hasDataAttr/hasFieldAttr/splitTrim/extractAllText/collectText/directText/hasContent/extractNonEmptyText
package stml

import (
	"testing"
)

func TestGetAttrAndHasAttr(t *testing.T) {
	n := firstElementNode(t, `<div id="x" data-fetch="/api"></div>`, "div")
	if got := getAttr(n, "id"); got != "x" {
		t.Errorf("getAttr id = %q, want x", got)
	}
	if got := getAttr(n, "missing"); got != "" {
		t.Errorf("getAttr missing = %q, want empty", got)
	}
	if !hasAttr(n, "data-fetch") {
		t.Errorf("hasAttr data-fetch should be true")
	}
	if hasAttr(n, "nope") {
		t.Errorf("hasAttr nope should be false")
	}
}
