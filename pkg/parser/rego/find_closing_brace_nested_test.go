//ff:func feature=policy type=test control=sequence
//ff:what findClosingBrace — 중첩된 `{}` 를 건너뛰고 외곽 `}` 위치 반환

package rego

import "testing"

func TestFindClosingBrace_Nested(t *testing.T) {
	s := "{inner}more}tail"
	if got := findClosingBrace(s); got != 11 {
		t.Errorf("got %d, want 11", got)
	}
}
