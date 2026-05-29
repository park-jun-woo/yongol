//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what bodyPropertyFromExpr 단위 테스트 (request.<prop> / request.body.<prop> 추출)

package ssac

import "testing"

func TestBodyPropertyFromExpr(t *testing.T) {
	cases := map[string]string{
		"request.payload":          "payload",
		"request.body.payload":     "payload",
		"  request.body.template ": "template",
		"request.body.a.b":         "a",
		"request.config more":      "config",
		"currentUser.ID":           "", // wrong prefix
		"":                         "",
		"request.":                 "",
	}
	for in, want := range cases {
		if got := bodyPropertyFromExpr(in); got != want {
			t.Errorf("bodyPropertyFromExpr(%q) = %q, want %q", in, got, want)
		}
	}
}
