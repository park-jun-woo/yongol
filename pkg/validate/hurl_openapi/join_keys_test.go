//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what joinKeys — 진단 메시지용 [a, b, c] 렌더링 검증

package hurl_openapi

import "testing"

func TestJoinKeys(t *testing.T) {
	cases := []struct {
		name string
		keys []string
		want string
	}{
		{name: "nil", keys: nil, want: "[]"},
		{name: "empty", keys: []string{}, want: "[]"},
		{name: "single", keys: []string{"a"}, want: "[a]"},
		{name: "two", keys: []string{"a", "b"}, want: "[a, b]"},
		{name: "three", keys: []string{"a", "b", "c"}, want: "[a, b, c]"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := joinKeys(c.keys)
			if got != c.want {
				t.Errorf("joinKeys(%v) = %q, want %q", c.keys, got, c.want)
			}
		})
	}
}
