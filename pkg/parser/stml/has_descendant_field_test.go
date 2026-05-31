//ff:func feature=stml-parse type=test control=sequence
//ff:what hasDescendantData/hasDescendantField/hasDescendantDataInFetch/extractParams/extractPageMetaAttrs
package stml

import (
	"testing"
)

func TestHasDescendantField(t *testing.T) {
	with := firstElementNode(t, `<div><p><span data-field="name"></span></p></div>`, "div")
	if !hasDescendantField(with) {
		t.Errorf("expected descendant data-field")
	}
	without := firstElementNode(t, `<div><p>text</p></div>`, "div")
	if hasDescendantField(without) {
		t.Errorf("expected no descendant field")
	}
}
