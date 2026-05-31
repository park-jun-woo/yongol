//ff:func feature=stml-parse type=test control=sequence
//ff:what isImplicitTag/isDataElement/getAttr/hasAttr/hasDataAttr/hasFieldAttr/splitTrim/extractAllText/collectText/directText/hasContent/extractNonEmptyText
package stml

import (
	"testing"
)

func TestDirectTextAndExtractNonEmptyText(t *testing.T) {
	n := firstElementNode(t, `<div>   <span>x</span>label</div>`, "div")
	// directText returns first non-empty direct text child; whitespace-only
	// text node is skipped, then the element, then "label".
	if got := directText(n); got != "label" {
		t.Errorf("directText = %q, want label", got)
	}
	// extractNonEmptyText on a non-text node returns ""
	span := firstElementNode(t, `<span>y</span>`, "span")
	if got := extractNonEmptyText(span); got != "" {
		t.Errorf("extractNonEmptyText(element) = %q, want empty", got)
	}
	if got := extractNonEmptyText(span.FirstChild); got != "y" {
		t.Errorf("extractNonEmptyText(text) = %q, want y", got)
	}
}
