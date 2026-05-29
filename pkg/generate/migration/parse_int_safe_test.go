//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestParseIntSafe — 숫자 문자열은 파싱, 비숫자 포함 시 0
package migration

import "testing"

func TestParseIntSafe(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"0", 0},
		{"42", 42},
		{"007", 7},
		{"", 0},
		{"12a", 0},
		{"-3", 0},
		{" 5", 0},
	}
	for _, c := range cases {
		if got := parseIntSafe(c.in); got != c.want {
			t.Errorf("parseIntSafe(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
