//ff:func feature=policy type=test control=sequence
//ff:what findClosingBrace — `}` 가 없으면 -1 반환

package rego

import "testing"

func TestFindClosingBrace_Unbalanced(t *testing.T) {
	s := "no closing"
	if got := findClosingBrace(s); got != -1 {
		t.Errorf("got %d, want -1", got)
	}
}
