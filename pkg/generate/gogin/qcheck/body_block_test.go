//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMaxBlockDepth — for/range/switch/typeswitch/select/case 각 노드 종류별 depth 검증
package qcheck

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func bodyBlock(t *testing.T, body string) *ast.BlockStmt {
	t.Helper()
	src := "package x\nfunc H() {\n" + body + "\n}"
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return file.Decls[0].(*ast.FuncDecl).Body
}
