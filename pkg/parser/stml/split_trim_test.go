//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what isImplicitTag/isDataElement/getAttr/hasAttr/hasDataAttr/hasFieldAttr/splitTrim/extractAllText/collectText/directText/hasContent/extractNonEmptyText
package stml

import (
	"testing"
)

func TestSplitTrim(t *testing.T) {
	got := splitTrim("a, b ,, c ")
	want := []string{"a", "b", "c"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
	if g := splitTrim("  ,  "); g != nil {
		t.Errorf("all-empty = %v, want nil", g)
	}
}
