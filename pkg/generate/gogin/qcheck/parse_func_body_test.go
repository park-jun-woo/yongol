//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestCollectResourceFromBlock — top-level + 중첩 if-블록 내 미닫힘 리소스 DF-06 재귀 수집
package qcheck

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func parseFuncBody(t *testing.T, src string) (*ast.BlockStmt, *token.FileSet) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "H" {
			return fn.Body, fset
		}
	}
	t.Fatalf("func H not found")
	return nil, nil
}
