//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestIsSelectorCall — pkg 매칭/임의 ident/비-selector/중첩 receiver/이름 불일치 분기 검증
package qcheck

import (
	"testing"
)

func TestIsSelectorCall(t *testing.T) {
	cases := []struct {
		expr, pkg, fn string
		want          bool
	}{
		{"json.Unmarshal(b, v)", "json", "Unmarshal", true},
		{"row.Scan(x)", "", "Scan", true},                 // empty pkg accepts any ident
		{"json.Unmarshal(b)", "yaml", "Unmarshal", false}, // wrong pkg
		{"json.Marshal(v)", "json", "Unmarshal", false},   // wrong func
		{"plain(x)", "", "plain", false},                  // not a selector
		{"a.b.Scan(x)", "", "Scan", false},                // receiver not a plain ident
	}
	for _, c := range cases {
		got := isSelectorCall(callExpr(t, c.expr), c.pkg, c.fn)
		if got != c.want {
			t.Errorf("isSelectorCall(%q, %q, %q) = %v, want %v", c.expr, c.pkg, c.fn, got, c.want)
		}
	}
}
