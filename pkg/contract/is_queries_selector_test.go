//ff:func feature=contract type=test control=sequence
//ff:what test: TestIsQueriesSelector — tail name "Queries" 일치/불일치/비-selector 분기 검증
package contract

import (
	"go/ast"
	"testing"
)

func TestIsQueriesSelector(t *testing.T) {
	// server.Queries — tail is "Queries"
	queries := &ast.SelectorExpr{
		X:   &ast.Ident{Name: "server"},
		Sel: &ast.Ident{Name: "Queries"},
	}
	// s.Other — tail is not "Queries"
	other := &ast.SelectorExpr{
		X:   &ast.Ident{Name: "s"},
		Sel: &ast.Ident{Name: "Other"},
	}
	// bare ident — not a selector at all
	bare := &ast.Ident{Name: "Queries"}

	if !isQueriesSelector(queries) {
		t.Errorf("expected true for server.Queries selector")
	}
	if isQueriesSelector(other) {
		t.Errorf("expected false for s.Other selector")
	}
	if isQueriesSelector(bare) {
		t.Errorf("expected false for non-selector expression")
	}
}
