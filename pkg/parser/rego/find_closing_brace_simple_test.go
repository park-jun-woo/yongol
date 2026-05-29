//ff:func feature=policy type=test control=sequence
//ff:what findClosingBrace — 깊이 0 에서 첫 `}` 위치 반환

package rego

import "testing"

func TestFindClosingBrace_Simple(t *testing.T) {
	s := "abc}rest"
	if got := findClosingBrace(s); got != 3 {
		t.Errorf("got %d, want 3", got)
	}
}
