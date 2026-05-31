//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"testing"
)

func TestNestifyPath_ZeroCov(t *testing.T) {
	cases := map[string]string{
		"/users/{id}":           "/users/:id",
		"/a/{x}/b/{y}":          "/a/:x/b/:y",
		"/static":               "/static",
		"/broken/{unterminated": "/broken/{unterminated",
	}
	for in, want := range cases {
		if got := nestifyPath(in); got != want {
			t.Errorf("nestifyPath(%q)=%q want %q", in, got, want)
		}
	}
}
