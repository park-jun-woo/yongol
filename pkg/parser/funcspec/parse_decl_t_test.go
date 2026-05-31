//ff:func feature=funcspec type=test control=sequence
//ff:what extractReturnTypes / isStubBody / processFuncDecl / findFuncDeclLine / extractGoParseErrorLine
package funcspec

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func parseDeclT(t *testing.T, src string) (*token.FileSet, *ast.File) {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return fset, f
}
