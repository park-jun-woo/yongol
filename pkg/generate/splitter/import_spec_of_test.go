//ff:func feature=gen-splitter type=test control=sequence
//ff:what primaryTypeName / sqlcSuffix / suffixFor / matchesOriginal / importName AST·도구 헬퍼
package splitter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func importSpecOf(t *testing.T, src string) *ast.ImportSpec {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	gd := f.Decls[0].(*ast.GenDecl)
	return gd.Specs[0].(*ast.ImportSpec)
}
