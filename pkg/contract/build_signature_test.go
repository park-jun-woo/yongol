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

func TestBuildSignatureWithError(t *testing.T) {
	sig := buildSigFromSrc(t, "func Do(id int64) (string, error) {}")
	if sig.Name != "Do" {
		t.Errorf("name: got %q want Do", sig.Name)
	}
	if len(sig.Params) != 1 || sig.Params[0] != (FuncParam{Name: "id", Type: "int64"}) {
		t.Errorf("params: got %+v", sig.Params)
	}
	if len(sig.Returns) != 2 || sig.Returns[0] != "string" || sig.Returns[1] != "error" {
		t.Errorf("returns: got %v", sig.Returns)
	}
	if !sig.HasErr {
		t.Errorf("expected HasErr true")
	}
}

func TestBuildSignatureNoReturns(t *testing.T) {
	sig := buildSigFromSrc(t, "func Run() {}")
	if sig.Name != "Run" {
		t.Errorf("name: got %q want Run", sig.Name)
	}
	if len(sig.Returns) != 0 {
		t.Errorf("returns: got %v want none", sig.Returns)
	}
	if sig.HasErr {
		t.Errorf("expected HasErr false")
	}
}

func TestBuildSignatureNonErrorReturn(t *testing.T) {
	sig := buildSigFromSrc(t, "func Count() int { return 0 }")
	if len(sig.Returns) != 1 || sig.Returns[0] != "int" {
		t.Errorf("returns: got %v want [int]", sig.Returns)
	}
	if sig.HasErr {
		t.Errorf("expected HasErr false for int return")
	}
}
