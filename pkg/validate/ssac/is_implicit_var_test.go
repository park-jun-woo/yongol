//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what isImplicitVar — 암시 소스 판정 (request/currentUser/query/message/other) 검증

package ssac

import "testing"

func TestIsImplicitVar(t *testing.T) {
	cases := []struct {
		name string
		want bool
	}{
		{"request", true},
		{"currentUser", true},
		{"query", true},
		{"message", true},
		{"user", false},
		{"", false},
		{"Request", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := isImplicitVar(c.name); got != c.want {
				t.Errorf("isImplicitVar(%q) = %v, want %v", c.name, got, c.want)
			}
		})
	}
}
