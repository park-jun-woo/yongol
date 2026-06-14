//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what normalizeTypeName — []/*/Wrapper[..]/pkg. 제거 후 베이스 타입명 추출

package openapi_ddl

import "testing"

func TestNormalizeTypeName(t *testing.T) {
	cases := []struct{ in, want string }{
		{"[]billing.CheckCreditsResponse", "CheckCreditsResponse"},
		{"*User", "User"},
		{"Page[Workflow]", "Workflow"},
		{"User", "User"},
		{"pkg.Thing", "Thing"},
		{"[]*pkg.Thing", "Thing"}, // strip [], strip *, then dot split
	}
	for _, c := range cases {
		if got := normalizeTypeName(c.in); got != c.want {
			t.Errorf("normalizeTypeName(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
