//ff:func feature=stml-parse type=test control=sequence
//ff:what isImplicitTag/isDataElement/getAttr/hasAttr/hasDataAttr/hasFieldAttr/splitTrim/extractAllText/collectText/directText/hasContent/extractNonEmptyText
package stml

import (
	"testing"
)

func TestHasContent(t *testing.T) {
	withText := firstElementNode(t, `<div>hello</div>`, "div")
	if !hasContent(withText) {
		t.Errorf("text content should count")
	}
	withChild := firstElementNode(t, `<div><span></span></div>`, "div")
	if !hasContent(withChild) {
		t.Errorf("element child should count")
	}
	empty := firstElementNode(t, `<div>   </div>`, "div")
	if hasContent(empty) {
		t.Errorf("whitespace-only should not count")
	}
}
