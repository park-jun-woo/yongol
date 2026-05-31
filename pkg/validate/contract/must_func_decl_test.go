//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what parse helpers — 테스트용 단일 expr/stmt/funcDecl AST 파싱 헬퍼
package contract

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func mustFuncDecl(t *testing.T, src string) *ast.FuncDecl {
	t.Helper()
	full := "package p\n" + src + "\n"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", full, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse func %q: %v", src, err)
	}
	for _, d := range f.Decls {
		if fd, ok := d.(*ast.FuncDecl); ok {
			return fd
		}
	}
	t.Fatalf("no func decl in %q", src)
	return nil
}
