//ff:func feature=gen-splitter type=test control=sequence
//ff:what keepImport / selectorPkgIdent / gatherSelectorNames / isSourceFile
package splitter

import (
	"go/ast"
	"testing"
)

func TestKeepImport(t *testing.T) {
	blank := &ast.ImportSpec{Name: &ast.Ident{Name: "_"}, Path: &ast.BasicLit{Value: `"x"`}}
	if !keepImport(blank, map[string]bool{}) {
		t.Errorf("blank import should be kept")
	}
	dot := &ast.ImportSpec{Name: &ast.Ident{Name: "."}, Path: &ast.BasicLit{Value: `"x"`}}
	if !keepImport(dot, map[string]bool{}) {
		t.Errorf("dot import should be kept")
	}
	used := &ast.ImportSpec{Path: &ast.BasicLit{Value: `"net/http"`}}
	if !keepImport(used, map[string]bool{"http": true}) {
		t.Errorf("used import should be kept")
	}
	if keepImport(used, map[string]bool{}) {
		t.Errorf("unused import should be dropped")
	}
}
