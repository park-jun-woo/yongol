//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestPySourceExpr — Go dot-access → Python dict key access 변환
package ssac

import (
	"testing"
)

func TestPySourceExpr(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"workflow", "workflow"},
		{"token.AccessToken", `token["access_token"]`},
		{"user.ID", `user["id"]`},
	}
	for _, c := range cases {
		if got := pySourceExpr(c.in); got != c.want {
			t.Errorf("pySourceExpr(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
