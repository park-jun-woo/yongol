//ff:func feature=rule type=test control=sequence
//ff:what TestStringSet_BasicOps — StringSet 기본 add/lookup 동작 검증

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
