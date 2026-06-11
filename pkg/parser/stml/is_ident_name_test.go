//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestIsIdentName — [A-Za-z_][A-Za-z0-9_]* 판정의 양성·음성(빈 문자열/숫자 선두/기호/비ASCII) 검증

package stml

import "testing"

func TestIsIdentName(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{"role", true},
		{"user_role", true},
		{"_r2", true},
		{"Role9", true},
		{"A", true},
		{"_", true},
		{"", false},
		{"9role", false}, // digit may not lead
		{"ro-le", false},
		{"ro.le", false},
		{"ro le", false},
		{"역할", false}, // ASCII only — emitted verbatim as a TS object key
	}
	for _, c := range cases {
		if got := isIdentName(c.s); got != c.want {
			t.Errorf("isIdentName(%q) = %v, want %v", c.s, got, c.want)
		}
	}
}
