//ff:func feature=gen-splitter type=test control=sequence
//ff:what primaryTypeName / sqlcSuffix / suffixFor / matchesOriginal / importName AST·도구 헬퍼
package splitter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func firstDecl(t *testing.T, src string) ast.Decl {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return f.Decls[0]
}
