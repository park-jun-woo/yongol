//ff:func feature=gen-splitter type=test control=sequence
//ff:what methodReceiver / receiverName / funcIdentifier / genDeclIdentifier / declIdentifier / funcFileName / genDeclFileName / fileNameForDecl
package splitter

import (
	"go/ast"
	"testing"
)

func TestMethodReceiverAndName(t *testing.T) {
	plain := declOf(t, "package p\nfunc Foo() {}").(*ast.FuncDecl)
	if got := methodReceiver(plain); got != "" {
		t.Errorf("plain func receiver = %q, want empty", got)
	}
	val := declOf(t, "package p\nfunc (q Queries) Bar() {}").(*ast.FuncDecl)
	if got := methodReceiver(val); got != "Queries" {
		t.Errorf("value receiver = %q, want Queries", got)
	}
	ptr := declOf(t, "package p\nfunc (q *Queries) Baz() {}").(*ast.FuncDecl)
	if got := methodReceiver(ptr); got != "Queries" {
		t.Errorf("pointer receiver = %q, want Queries", got)
	}
}
