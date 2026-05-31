//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestCollectScanFromBlock — 중첩 블록 내 미가드 .Scan() DF-02 재귀 수집 + 가드 케이스
package qcheck

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func parseHBody(t *testing.T, src string) (*ast.BlockStmt, *token.FileSet) {
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
