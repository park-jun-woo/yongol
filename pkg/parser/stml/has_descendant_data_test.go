//ff:func feature=stml-parse type=test control=sequence
//ff:what hasDescendantData/hasDescendantField/hasDescendantDataInFetch/extractParams/extractPageMetaAttrs
package stml

import (
	"testing"
)

func TestHasDescendantData(t *testing.T) {
	with := firstElementNode(t, `<div><section><span data-fetch="/a"></span></section></div>`, "div")
	if !hasDescendantData(with) {
		t.Errorf("expected descendant data-fetch")
	}
	without := firstElementNode(t, `<div><section><span class="c"></span></section></div>`, "div")
	if hasDescendantData(without) {
		t.Errorf("expected no descendant data")
	}
}
