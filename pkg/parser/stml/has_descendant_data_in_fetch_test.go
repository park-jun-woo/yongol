//ff:func feature=stml-parse type=test control=sequence
//ff:what hasDescendantData/hasDescendantField/hasDescendantDataInFetch/extractParams/extractPageMetaAttrs
package stml

import (
	"testing"
)

func TestHasDescendantDataInFetch(t *testing.T) {
	with := firstElementNode(t, `<div><ul><li data-field="x"></li></ul></div>`, "div")
	if !hasDescendantDataInFetch(with) {
		t.Errorf("expected descendant data-* attr")
	}
	without := firstElementNode(t, `<div><ul><li class="c"></li></ul></div>`, "div")
	if hasDescendantDataInFetch(without) {
		t.Errorf("expected no descendant data-* attr")
	}
}
