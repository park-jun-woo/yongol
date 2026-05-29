//ff:func feature=agent type=test control=sequence
//ff:what TestIsImmutable — features.yaml/.hurl/.yongol 만 immutable 로 판별 검증

package agent

import "testing"

func TestIsImmutable(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"specs/features.yaml", true},
		{"a/b/tests/login.hurl", true},
		{"specs/.yongol", true},
		{"specs/openapi.yaml", false},
		{"db/queries/user.sql", false},
		{"", false},
	}
	for _, c := range cases {
		if got := isImmutable(c.in); got != c.want {
			t.Errorf("isImmutable(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
