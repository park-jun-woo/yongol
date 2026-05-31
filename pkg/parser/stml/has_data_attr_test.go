//ff:func feature=stml-parse type=test control=sequence
//ff:what isImplicitTag/isDataElement/getAttr/hasAttr/hasDataAttr/hasFieldAttr/splitTrim/extractAllText/collectText/directText/hasContent/extractNonEmptyText
package stml

import (
	"testing"
)

func TestHasDataAttr(t *testing.T) {
	withData := firstElementNode(t, `<div data-x="1"></div>`, "div")
	if !hasDataAttr(withData) {
		t.Errorf("expected data attr")
	}
	without := firstElementNode(t, `<div class="c"></div>`, "div")
	if hasDataAttr(without) {
		t.Errorf("expected no data attr")
	}
}
