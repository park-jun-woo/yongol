//ff:func feature=gen-splitter type=test control=sequence
//ff:what docOf / funcDoc / genDeclDoc / detectControl / controlFor / funcTypeFor / extractHeader
package splitter

import (
	"go/ast"
	"testing"
)

func TestFuncDoc(t *testing.T) {
	d := parseDeclWithComments(t, "package p\n// Foo does things.\nfunc Foo() {}").(*ast.FuncDecl)
	if got := funcDoc(d); got != "Foo does things.\n" {
		t.Errorf("funcDoc = %q", got)
	}
	nodoc := parseDeclWithComments(t, "package p\nfunc Bar() {}").(*ast.FuncDecl)
	if got := funcDoc(nodoc); got != "" {
		t.Errorf("no-doc funcDoc = %q, want empty", got)
	}
}
