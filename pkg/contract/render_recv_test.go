//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestRenderRecv — Ident 수신자는 이름 직반환, 그 외는 printer 문자열화 검증

package contract

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestRenderRecv(t *testing.T) {
	fset := token.NewFileSet()

	// Simple Ident receiver returns its name directly.
	if got := renderRecv(fset, &ast.Ident{Name: "user"}); got != "user" {
		t.Errorf("ident: got %q want user", got)
	}

	// Non-Ident receiver (a selector chain) is rendered via printer.
	expr, err := parser.ParseExpr("a.b")
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if got := renderRecv(fset, expr); got != "a.b" {
		t.Errorf("selector: got %q want a.b", got)
	}

	// Index expression also routes through printer.
	idxExpr, err := parser.ParseExpr("arr[0]")
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if got := renderRecv(fset, idxExpr); got != "arr[0]" {
		t.Errorf("index: got %q want arr[0]", got)
	}
}
