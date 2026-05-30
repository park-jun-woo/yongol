//ff:func feature=cli-init type=test control=iteration dimension=1
//ff:what TestCollapseUnderscores — 다중 언더스코어 축약 + edge trim 검증

package cliinit

import "testing"

func TestCollapseUnderscores(t *testing.T) {
	cases := []struct{ in, want string }{
		{"_Foo__Bar_", "Foo_Bar"},
		{"a___b", "a_b"},
		{"already_clean", "already_clean"},
		{"___", ""},
		{"", ""},
		{"no_underscore_runs", "no_underscore_runs"},
	}
	for _, c := range cases {
		if got := collapseUnderscores(c.in); got != c.want {
			t.Errorf("collapseUnderscores(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
