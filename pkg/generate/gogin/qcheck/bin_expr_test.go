//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBinaryIsErrNilCheck — err!=nil / nil!=err / 비-ident / 무-err 분기 검증
package qcheck

import (
	"go/ast"
	"go/parser"
	"testing"
)

func binExpr(t *testing.T, expr string) *ast.BinaryExpr {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatalf("parse %q: %v", expr, err)
	}
	bin, ok := e.(*ast.BinaryExpr)
	if !ok {
		t.Fatalf("expr %q is not BinaryExpr: %T", expr, e)
	}
	return bin
}
