//ff:func feature=funcspec type=test control=sequence
//ff:what TestFuncspecHelpers — unit tests for the pure funcspec parser helper functions
package funcspec

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func parseBody(t *testing.T, body string) *ast.BlockStmt {
	t.Helper()
	src := "package p\nfunc f() (Resp, error) {\n" + body + "\n}\n"
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, 0)
	if err != nil {
		t.Fatalf("parse body: %v", err)
	}
	return file.Decls[0].(*ast.FuncDecl).Body
}
