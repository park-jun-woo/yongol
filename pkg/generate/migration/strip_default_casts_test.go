//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestStripDefaultCasts — DEFAULT 표현식 끝 ::type 캐스트 반복 제거
package migration

import "testing"

func TestStripDefaultCasts(t *testing.T) {
	cases := []struct{ in, want string }{
		{"'active'::text", "'active'"},
		{"0::integer", "0"},
		{"'x'::character varying", "'x'"},
		{"42", "42"},
		{"now()", "now()"},
	}
	for _, c := range cases {
		if got := stripDefaultCasts(c.in); got != c.want {
			t.Errorf("stripDefaultCasts(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
