//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestSSaCParseHelpers — unit tests for the pure ssac parser helper functions
package ssac

import (
	"go/ast"
	"go/parser"
	"testing"
)

func TestExprToString(t *testing.T) {
	mustExpr := func(src string) ast.Expr {
		e, err := parser.ParseExpr(src)
		if err != nil {
			t.Fatalf("parse %q: %v", src, err)
		}
		return e
	}
	if got := exprToString(mustExpr("foo")); got != "foo" {
		t.Errorf("ident = %q", got)
	}
	if got := exprToString(mustExpr("pkg.Type")); got != "pkg.Type" {
		t.Errorf("selector = %q", got)
	}
	if got := exprToString(mustExpr("a.b.c")); got != "a.b.c" {
		t.Errorf("nested selector = %q", got)
	}
	// Unsupported expr → "".
	if got := exprToString(mustExpr("[]int")); got != "" {
		t.Errorf("array expr → %q, want empty", got)
	}
}
