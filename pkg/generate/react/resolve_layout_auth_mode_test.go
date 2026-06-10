//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what resolveLayoutAuthMode — auth 게이트 3분기 ("", bearer, cookie) 매핑 검증

package react

import "testing"

func TestResolveLayoutAuthMode(t *testing.T) {
	cases := []struct {
		name                string
		hasAuth, bearerAuth bool
		want                string
	}{
		{"no auth", false, false, ""},
		{"bearer", true, true, "bearer"},
		{"cookie or hybrid", true, false, "cookie"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := resolveLayoutAuthMode(c.hasAuth, c.bearerAuth); got != c.want {
				t.Errorf("resolveLayoutAuthMode(%v, %v) = %q, want %q", c.hasAuth, c.bearerAuth, got, c.want)
			}
		})
	}
}
