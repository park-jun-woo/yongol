//ff:func feature=rule type=test control=sequence
//ff:what TestStringSet_NilLookupReturnsFalse — nil StringSet 의 lookup 이 panic 없이 false 를 반환

package rule

import "testing"

func TestStringSet_NilLookupReturnsFalse(t *testing.T) {
	var s StringSet
	// Reading from a nil map returns the zero value — does not panic.
	if s["anything"] {
		t.Fatalf("nil StringSet lookup = true; want false")
	}
}
