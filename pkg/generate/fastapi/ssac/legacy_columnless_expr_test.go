//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestLegacyColumnlessExpr — legacyColumnlessExpr source→Python 식별자 매핑 검증
package ssac

import "testing"

func TestLegacyColumnlessExpr(t *testing.T) {
	cases := map[string]string{
		"request":     "params",
		"currentUser": "current_user",
		"foo":         "foo",
		"":            "",
	}
	for in, want := range cases {
		if got := legacyColumnlessExpr(in); got != want {
			t.Errorf("legacyColumnlessExpr(%q) = %q, want %q", in, got, want)
		}
	}
}
