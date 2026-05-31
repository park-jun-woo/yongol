//ff:func feature=gen-splitter type=test control=sequence
//ff:what docOf / funcDoc / genDeclDoc / detectControl / controlFor / funcTypeFor / extractHeader
package splitter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func parseDeclWithComments(t *testing.T, src string) ast.Decl {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return f.Decls[0]
}
