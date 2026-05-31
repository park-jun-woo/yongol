//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestAssignCallsAndGuarded — guarded shape true + 각 거부 분기 검증
package qcheck

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func blockStmts(t *testing.T, body string) []ast.Stmt {
	t.Helper()
	src := "package x\nfunc f() {\n" + body + "\n}"
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	fn := file.Decls[0].(*ast.FuncDecl)
	return fn.Body.List
}
