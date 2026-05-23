//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-statemachine
//ff:what normPath — path 세그먼트 정규화 검증

package hurl_statemachine

import "testing"

func TestNormPath(t *testing.T) {
	cases := []struct {
		name string
		path string
		want string
	}{
		{name: "simple", path: "/users", want: "/users"},
		{name: "openapi_var", path: "/users/{id}", want: "/users/:param"},
		{name: "hurl_var", path: "/users/{{id}}", want: "/users/:param"},
		{name: "numeric", path: "/users/123", want: "/users/:param"},
		{name: "mixed", path: "/users/{id}/posts/42", want: "/users/:param/posts/:param"},
		{name: "trailing_slash", path: "/users/", want: "/users"},
		{name: "root", path: "/", want: "/"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := normPath(c.path, reOpenAPIVarKey, nil)
			if got != c.want {
				t.Errorf("normPath(%q) = %q, want %q", c.path, got, c.want)
			}
		})
	}
}
