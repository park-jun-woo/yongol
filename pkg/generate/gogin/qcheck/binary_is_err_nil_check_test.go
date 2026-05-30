//ff:func feature=gen-gogin type=test control=branch topic=err-guard
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

func TestBinaryIsErrNilCheck(t *testing.T) {
	cases := []struct {
		expr string
		want bool
	}{
		{"err != nil", true},
		{"nil != err", true},
		{"myErr != nil", true},
		{"x != nil", false},
		{"nil != x", false},
		{"nil != nil", false},
		{"err != other", false}, // neither side is nil
		{"foo() != nil", false}, // LHS not an Ident
	}
	for _, c := range cases {
		if got := binaryIsErrNilCheck(binExpr(t, c.expr)); got != c.want {
			t.Errorf("binaryIsErrNilCheck(%q) = %v, want %v", c.expr, got, c.want)
		}
	}
}
