//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestExtractRuleID — 대괄호/콜론 형식의 룰 ID 추출 검증
package agent

import (
	"testing"
)

func TestExtractRuleID(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"bracket", "[S-74] something wrong", "S-74"},
		{"colon", "S-74: something wrong", "S-74"},
		{"no id", "just a message without id markers here", ""},
		{"empty", "", ""},
		{"colon too far", "this is a very long prefix before any colon: x", ""},
	}
	for _, c := range cases {
		if got := extractRuleID(c.in); got != c.want {
			t.Errorf("%s: extractRuleID(%q) = %q, want %q", c.name, c.in, got, c.want)
		}
	}
}
