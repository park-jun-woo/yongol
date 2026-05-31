//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what isImplicitTag/isDataElement/getAttr/hasAttr/hasDataAttr/hasFieldAttr/splitTrim/extractAllText/collectText/directText/hasContent/extractNonEmptyText
package stml

import (
	"testing"
)

func TestIsImplicitTag(t *testing.T) {
	for _, tag := range []string{"html", "head", "body"} {
		if !isImplicitTag(tag) {
			t.Errorf("%q should be implicit", tag)
		}
	}
	if isImplicitTag("div") {
		t.Errorf("div should not be implicit")
	}
}
