//ff:func feature=gen-splitter type=test control=sequence
//ff:what methodReceiver / receiverName / funcIdentifier / genDeclIdentifier / declIdentifier / funcFileName / genDeclFileName / fileNameForDecl
package splitter

import (
	"go/ast"
	"testing"
)

func TestFuncIdentifier(t *testing.T) {
	plain := declOf(t, "package p\nfunc Foo() {}").(*ast.FuncDecl)
	if got := funcIdentifier(plain); got != "Foo" {
		t.Errorf("got %q, want Foo", got)
	}
	method := declOf(t, "package p\nfunc (q *Queries) FindUser() {}").(*ast.FuncDecl)
	if got := funcIdentifier(method); got != "Queries.FindUser" {
		t.Errorf("got %q, want Queries.FindUser", got)
	}
}
