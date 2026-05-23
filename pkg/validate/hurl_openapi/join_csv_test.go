//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what joinCSV — comma-separated join 검증

package hurl_openapi

import "testing"

func TestJoinCSV(t *testing.T) {
	cases := []struct {
		name string
		s    []string
		want string
	}{
		{name: "nil", s: nil, want: ""},
		{name: "empty", s: []string{}, want: ""},
		{name: "single", s: []string{"a"}, want: "a"},
		{name: "two", s: []string{"a", "b"}, want: "a, b"},
		{name: "three", s: []string{"a", "b", "c"}, want: "a, b, c"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := joinCSV(c.s)
			if got != c.want {
				t.Errorf("joinCSV(%v) = %q, want %q", c.s, got, c.want)
			}
		})
	}
}
