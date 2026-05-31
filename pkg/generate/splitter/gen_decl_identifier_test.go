//ff:func feature=gen-splitter type=test control=sequence
//ff:what methodReceiver / receiverName / funcIdentifier / genDeclIdentifier / declIdentifier / funcFileName / genDeclFileName / fileNameForDecl
package splitter

import (
	"go/ast"
	"testing"
)

func TestGenDeclIdentifier(t *testing.T) {
	typ := declOf(t, "package p\ntype Row struct{}").(*ast.GenDecl)
	if got := genDeclIdentifier(typ); got != "Row" {
		t.Errorf("got %q, want Row", got)
	}
	multi := declOf(t, "package p\nconst (\nA=1\nB=2\n)").(*ast.GenDecl)
	if got := genDeclIdentifier(multi); got != "const" {
		t.Errorf("got %q, want const", got)
	}
}
