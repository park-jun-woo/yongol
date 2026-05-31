//ff:func feature=funcspec type=test control=sequence
//ff:what TestFuncspecHelpers — unit tests for the pure funcspec parser helper functions
package funcspec

import (
	"go/ast"
	"testing"
)

func TestAllZeroResults(t *testing.T) {
	if !allZeroResults([]ast.Expr{parseExpr(t, "nil"), parseExpr(t, "Resp{}")}) {
		t.Error("all-zero results should be true")
	}
	if allZeroResults([]ast.Expr{parseExpr(t, "nil"), parseExpr(t, "42")}) {
		t.Error("mixed results should be false")
	}
	// empty list → true.
	if !allZeroResults(nil) {
		t.Error("empty results should be true")
	}
}
