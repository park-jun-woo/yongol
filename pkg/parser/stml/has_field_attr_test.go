//ff:func feature=stml-parse type=test control=sequence
//ff:what isImplicitTag/isDataElement/getAttr/hasAttr/hasDataAttr/hasFieldAttr/splitTrim/extractAllText/collectText/directText/hasContent/extractNonEmptyText
package stml

import (
	"testing"
)

func TestHasFieldAttr(t *testing.T) {
	field := firstElementNode(t, `<span data-field="name"></span>`, "span")
	if !hasFieldAttr(field) {
		t.Errorf("data-field should match")
	}
	none := firstElementNode(t, `<span data-component="c"></span>`, "span")
	if hasFieldAttr(none) {
		t.Errorf("data-component alone should not match")
	}
}
