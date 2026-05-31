//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestBuildSignature — FuncDecl→FuncSignature 변환, error 반환 시 HasErr, 반환 없음 분기 검증
package contract

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func buildSigFromSrc(t *testing.T, src string) FuncSignature {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", "package p\n"+src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	for _, decl := range f.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok {
			return buildSignature(fset, fd)
		}
	}
	t.Fatalf("no func decl in %q", src)
	return FuncSignature{}
}
