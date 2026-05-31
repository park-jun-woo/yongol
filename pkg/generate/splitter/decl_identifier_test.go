//ff:func feature=gen-splitter type=test control=sequence
//ff:what methodReceiver / receiverName / funcIdentifier / genDeclIdentifier / declIdentifier / funcFileName / genDeclFileName / fileNameForDecl
package splitter

import (
	"testing"
)

func TestDeclIdentifier(t *testing.T) {
	fn := declOf(t, "package p\nfunc Foo() {}")
	if got := declIdentifier(fn); got != "Foo" {
		t.Errorf("func got %q", got)
	}
	gd := declOf(t, "package p\ntype Row struct{}")
	if got := declIdentifier(gd); got != "Row" {
		t.Errorf("gendecl got %q", got)
	}
}
