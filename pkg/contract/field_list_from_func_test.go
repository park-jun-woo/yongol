//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestExpandFieldList — nil 리스트, 다중 이름 그룹, 익명 반환 타입 전개 분기 검증
package contract

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func fieldListFromFunc(t *testing.T, src string) (*token.FileSet, *ast.FuncDecl) {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", "package p\n"+src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	for _, decl := range f.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok {
			return fset, fd
		}
	}
	t.Fatalf("no func decl in %q", src)
	return nil, nil
}
