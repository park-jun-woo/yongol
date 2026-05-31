//ff:func feature=validate type=test control=sequence
//ff:what TestContainsStr_ZeroCov — containsStr 포함/미포함 분기 직접 호출

package features

import "testing"

func TestContainsStr_ZeroCov(t *testing.T) {
	s := []string{"a", "b", "c"}
	if !containsStr(s, "b") {
		t.Errorf("containsStr should find b")
	}
	if containsStr(s, "z") {
		t.Errorf("containsStr should not find z")
	}
	if containsStr(nil, "x") {
		t.Errorf("nil slice should be false")
	}
}
