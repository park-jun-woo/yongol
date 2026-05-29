//ff:func feature=ground type=test control=iteration dimension=1
//ff:what capitalizeFirst — 첫 글자 대문자 변환 + non-letter 보존

package ground

import "testing"

func TestCapitalizeFirst_Cases(t *testing.T) {
	cases := []struct{ in, want string }{
		{"", ""},
		{"hashPassword", "HashPassword"},
		{"a", "A"},
		{"Z", "Z"},       // already upper
		{"_x", "_x"},     // non-letter unchanged
		{"1foo", "1foo"}, // digit unchanged
	}
	for _, c := range cases {
		if got := capitalizeFirst(c.in); got != c.want {
			t.Errorf("capitalizeFirst(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
