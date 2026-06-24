//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestDomainIdentTitle — domainIdent/domainTitle 정규화 분기 검증 (BUG-142)

package middleware

import "testing"

func TestDomainIdentTitle(t *testing.T) {
	cases := []struct{ in, ident, title string }{
		{"public", "public", "Public"},
		{"my-admin", "my_admin", "MyAdmin"},
		{"2nd", "_2nd", "2nd"},
		{"", "_", ""},
		{"ABC", "abc", "Abc"},
	}
	for _, c := range cases {
		if got := domainIdent(c.in); got != c.ident {
			t.Errorf("domainIdent(%q)=%q want %q", c.in, got, c.ident)
		}
		if got := domainTitle(c.in); got != c.title {
			t.Errorf("domainTitle(%q)=%q want %q", c.in, got, c.title)
		}
	}
}
