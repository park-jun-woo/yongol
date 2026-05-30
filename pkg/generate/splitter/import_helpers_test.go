//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what keepImport / selectorPkgIdent / gatherSelectorNames / isSourceFile

package splitter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
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

func TestGatherSelectorNames(t *testing.T) {
	src := "package p\nfunc f() { http.Get(\"u\"); fmt.Println(1) }"
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "x.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	used := map[string]bool{}
	gatherSelectorNames(file.Decls[0], used)
	if !used["http"] || !used["fmt"] {
		t.Errorf("expected http and fmt collected, got %v", used)
	}
}

func TestIsSourceFile(t *testing.T) {
	dir := t.TempDir()
	// plain source matching pattern with no ff header -> true
	src := filepath.Join(dir, "api.gen.go")
	if err := os.WriteFile(src, []byte("package p\nfunc f(){}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if !isSourceFile(src, ToolOAPICodegen) {
		t.Errorf("plain .gen.go should be a source file")
	}
	// already-split output (ff:func header) -> false
	split := filepath.Join(dir, "f.gen.go")
	if err := os.WriteFile(split, []byte("//ff:func x\npackage p\nfunc f(){}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if isSourceFile(split, ToolOAPICodegen) {
		t.Errorf("ff-annotated output should not be a source file")
	}
	// name not matching tool pattern -> false
	other := filepath.Join(dir, "plain.go")
	if err := os.WriteFile(other, []byte("package p\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if isSourceFile(other, ToolOAPICodegen) {
		t.Errorf("non-matching name should not be a source file")
	}
}
