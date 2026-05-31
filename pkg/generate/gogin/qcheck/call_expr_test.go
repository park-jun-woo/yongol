//ff:func feature=gen-gogin type=test control=sequence
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
