//ff:func feature=gen-gogin type=test control=branch topic=err-guard
//ff:what TestIsSelectorCall — pkg 매칭/임의 ident/비-selector/중첩 receiver/이름 불일치 분기 검증

package qcheck

import (
	"go/ast"
	"go/parser"
	"testing"
)

func callExpr(t *testing.T, expr string) *ast.CallExpr {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatalf("parse %q: %v", expr, err)
	}
	call, ok := e.(*ast.CallExpr)
	if !ok {
		t.Fatalf("expr %q is not CallExpr: %T", expr, e)
	}
	return call
}

func TestIsSelectorCall(t *testing.T) {
	cases := []struct {
		expr, pkg, fn string
		want          bool
	}{
		{"json.Unmarshal(b, v)", "json", "Unmarshal", true},
		{"row.Scan(x)", "", "Scan", true},      // empty pkg accepts any ident
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
