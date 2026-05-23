//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what quoted — 문자열을 쌍따옴표로 감싸는 검증

package manifest

import "testing"

func TestQuoted(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"hello", "\"hello\""},
		{"", "\"\""},
		{"a b", "\"a b\""},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			got := quoted(c.in)
			if got != c.want {
				t.Errorf("quoted(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
