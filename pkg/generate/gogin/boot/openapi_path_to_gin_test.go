//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=path-convert
//ff:what openAPIPathToGin — OpenAPI "/x/{id}" → gin "/x/:id" 경로 변환

package boot

import "testing"

func TestOpenAPIPathToGin(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"single param", "/users/{id}", "/users/:id"},
		{"two params", "/orgs/{org}/repos/{repo}", "/orgs/:org/repos/:repo"},
		{"no params", "/health", "/health"},
		{"root", "/", "/"},
		{"empty", "", ""},
		{"unterminated brace stays literal", "/users/{id", "/users/{id"},
		{"trailing param", "/a/{b}/", "/a/:b/"},
	}
	for _, c := range cases {
		if got := openAPIPathToGin(c.in); got != c.want {
			t.Errorf("%s: openAPIPathToGin(%q) = %q, want %q", c.name, c.in, got, c.want)
		}
	}
}
