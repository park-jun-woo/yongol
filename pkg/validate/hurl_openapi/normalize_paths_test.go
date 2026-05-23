//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what normalizeHurlPath — Hurl 요청 path를 세그먼트 배열로 정규화 검증

package hurl_openapi

import "testing"

func TestNormalizeHurlPath(t *testing.T) {
	cases := []struct {
		name string
		path string
		want []string
	}{
		{name: "simple", path: "/users", want: []string{"users"}},
		{name: "nested", path: "/users/posts", want: []string{"users", "posts"}},
		{name: "hurl_var", path: "/users/{{id}}", want: []string{"users", ":param"}},
		{name: "numeric_literal", path: "/users/123", want: []string{"users", ":param"}},
		{name: "query_stripped", path: "/users?page=1", want: []string{"users"}},
		{name: "trailing_slash", path: "/users/", want: []string{"users"}},
		{name: "mixed_segments", path: "/users/{{id}}/posts/42", want: []string{"users", ":param", "posts", ":param"}},
		{name: "root", path: "/", want: nil},
		{name: "empty", path: "", want: nil},
		{name: "whitespace_trimmed", path: "  /users  ", want: []string{"users"}},
		{name: "literal_text_kept", path: "/api/v1/users", want: []string{"api", "v1", "users"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runStringSliceCase(t, normalizeHurlPath(c.path), c.want)
		})
	}
}
