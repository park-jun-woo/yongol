//ff:func feature=migration type=test control=selection dimension=5
//ff:what TestLooksLikeCastTarget — 단순 식별자/quoted/array 는 true, 비식별자는 false
package migration

import "testing"

func TestLooksLikeCastTarget(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"text", true},
		{`"MyType"`, true},
		{"integer[]", true},
		{"character varying", true},
		{"", false},
		{"now()", false},
		{"a+b", false},
	}
	for _, c := range cases {
		if got := looksLikeCastTarget(c.in); got != c.want {
			t.Errorf("looksLikeCastTarget(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
