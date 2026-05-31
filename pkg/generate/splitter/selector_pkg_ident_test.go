//ff:func feature=gen-splitter type=test control=sequence
//ff:what keepImport / selectorPkgIdent / gatherSelectorNames / isSourceFile
package splitter

import (
	"go/ast"
	"testing"
)

func TestSelectorPkgIdent(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "http"}, Sel: &ast.Ident{Name: "Get"}}
	if id := selectorPkgIdent(sel); id == nil || id.Name != "http" {
		t.Errorf("expected http ident, got %v", id)
	}
	// chained selector: X is itself a SelectorExpr -> nil
	chained := &ast.SelectorExpr{X: sel, Sel: &ast.Ident{Name: "Do"}}
	if id := selectorPkgIdent(chained); id != nil {
		t.Errorf("chained selector should yield nil, got %v", id)
	}
	if id := selectorPkgIdent(&ast.Ident{Name: "x"}); id != nil {
		t.Errorf("non-selector should yield nil")
	}
}
