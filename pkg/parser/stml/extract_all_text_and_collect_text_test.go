//ff:func feature=stml-parse type=test control=sequence
//ff:what isImplicitTag/isDataElement/getAttr/hasAttr/hasDataAttr/hasFieldAttr/splitTrim/extractAllText/collectText/directText/hasContent/extractNonEmptyText
package stml

import (
	"strings"
	"testing"
)

func TestExtractAllTextAndCollectText(t *testing.T) {
	n := firstElementNode(t, `<div>Hello <b>World</b>!</div>`, "div")
	if got := extractAllText(n); got != "Hello World!" {
		t.Errorf("extractAllText = %q, want %q", got, "Hello World!")
	}
	var sb strings.Builder
	collectText(n, &sb)
	if !strings.Contains(sb.String(), "World") {
		t.Errorf("collectText missing World: %q", sb.String())
	}
}
