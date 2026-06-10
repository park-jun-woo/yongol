//ff:func feature=stml-parse type=test control=sequence
//ff:what hasDescendantOnError — 하위 data-on-error 존재/부재 검증

package stml

import (
	"testing"
)

func TestHasDescendantOnError(t *testing.T) {
	with := firstElementNode(t, `<div><section><span data-on-error="msg"></span></section></div>`, "div")
	if !hasDescendantOnError(with) {
		t.Errorf("expected descendant data-on-error")
	}
	without := firstElementNode(t, `<div><section><span class="c"></span></section></div>`, "div")
	if hasDescendantOnError(without) {
		t.Errorf("expected no descendant data-on-error")
	}
}
