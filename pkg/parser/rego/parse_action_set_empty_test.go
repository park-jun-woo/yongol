//ff:func feature=policy type=test control=sequence
//ff:what parseActionSet — 빈 문자열은 nil 반환

package rego

import "testing"

func TestParseActionSet_Empty(t *testing.T) {
	if got := parseActionSet(""); got != nil {
		t.Errorf("got %v, want nil", got)
	}
}
