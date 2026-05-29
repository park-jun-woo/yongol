//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestItoa — 정수→문자열 변환 (음수/0/양수)
package migration

import "testing"

func TestItoa(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{0, "0"},
		{42, "42"},
		{-7, "-7"},
	}
	for _, c := range cases {
		if got := itoa(c.in); got != c.want {
			t.Errorf("itoa(%d) = %q, want %q", c.in, got, c.want)
		}
	}
}
