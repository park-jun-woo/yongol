//ff:func feature=migration type=test control=sequence
//ff:what TestNormalizeIndexMethod — "" 와 btree 는 동일 토큰, 그 외는 보존
package migration

import "testing"

func TestNormalizeIndexMethod(t *testing.T) {
	cases := []struct{ in, want string }{
		{"", "btree"},
		{"btree", "btree"},
		{"gin", "gin"},
		{"gist", "gist"},
	}
	for _, c := range cases {
		if got := normalizeIndexMethod(c.in); got != c.want {
			t.Errorf("normalizeIndexMethod(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
