//ff:func feature=stml-parse type=test control=sequence
//ff:what isImplicitTag/isDataElement/getAttr/hasAttr/hasDataAttr/hasFieldAttr/splitTrim/extractAllText/collectText/directText/hasContent/extractNonEmptyText
package stml

import (
	"testing"
)

func TestIsDataElement(t *testing.T) {
	fetch := firstElementNode(t, `<div data-fetch="/a"></div>`, "div")
	if !isDataElement(fetch) {
		t.Errorf("data-fetch should be data element")
	}
	action := firstElementNode(t, `<form data-action="/a"></form>`, "form")
	if !isDataElement(action) {
		t.Errorf("data-action should be data element")
	}
	plain := firstElementNode(t, `<div class="c"></div>`, "div")
	if isDataElement(plain) {
		t.Errorf("plain div should not be data element")
	}
}
