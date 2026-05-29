//ff:func feature=migration type=test control=selection dimension=4
//ff:what TestCanonIdent — quoted 보존, 그 외 trim+소문자화+후행 콤마 제거
package migration

import "testing"

func TestCanonIdent(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"  Users  ", "users"},
		{"id,", "id"},
		{`"MixedCase"`, "MixedCase"},
		{"EMAIL", "email"},
	}
	for _, c := range cases {
		if got := canonIdent(c.in); got != c.want {
			t.Errorf("canonIdent(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
