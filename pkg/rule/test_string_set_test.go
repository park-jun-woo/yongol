//ff:func feature=rule type=test control=sequence
//ff:what StringSet — map[string]bool 별칭 — 기본 lookup/add/zero-value 동작 검증

package rule

import "testing"

func TestStringSet_BasicOps(t *testing.T) {
	s := StringSet{}
	s["a"] = true
	if !s["a"] {
		t.Fatalf("StringSet[a] = false after add; want true")
	}
	if s["b"] {
		t.Fatalf("StringSet[b] = true without add; want false (zero value)")
	}
	s["b"] = true
	if !s["b"] {
		t.Fatalf("StringSet[b] = false after add; want true")
	}
}

func TestStringSet_NilLookupReturnsFalse(t *testing.T) {
	var s StringSet
	// Reading from a nil map returns the zero value — does not panic.
	if s["anything"] {
		t.Fatalf("nil StringSet lookup = true; want false")
	}
}
