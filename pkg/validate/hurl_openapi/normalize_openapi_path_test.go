//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what normalizeOpenAPIPath — OpenAPI path를 세그먼트 배열로 정규화 검증

package hurl_openapi

import "testing"

func TestNormalizeOpenAPIPath(t *testing.T) {
	cases := []struct {
		name string
		path string
		want []string
	}{
		{name: "simple", path: "/users", want: []string{"users"}},
		{name: "nested", path: "/users/posts", want: []string{"users", "posts"}},
		{name: "param", path: "/users/{id}", want: []string{"users", ":param"}},
		{name: "multiple_params", path: "/users/{id}/posts/{postId}", want: []string{"users", ":param", "posts", ":param"}},
		{name: "trailing_slash", path: "/users/", want: []string{"users"}},
		{name: "leading_trailing_slash", path: "/users/posts/", want: []string{"users", "posts"}},
		{name: "root", path: "/", want: nil},
		{name: "empty", path: "", want: nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runStringSliceCase(t, normalizeOpenAPIPath(c.path), c.want)
		})
	}
}
