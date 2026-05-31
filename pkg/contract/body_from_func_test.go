//ff:func feature=contract type=test control=sequence
//ff:what test: TestCollectCallSelector — Queries 메서드/패키지 호출/denylist/비-exported 분기를 queries·calls 맵에 분류 검증
package contract

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func bodyFromFunc(t *testing.T, body string) (*token.FileSet, *ast.BlockStmt) {
	t.Helper()
	src := "package p\nfunc F() {\n" + body + "\n}\n"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	fd := f.Decls[0].(*ast.FuncDecl)
	return fset, fd.Body
}
