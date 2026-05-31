//ff:func feature=gen-splitter type=test control=sequence
//ff:what docOf / funcDoc / genDeclDoc / detectControl / controlFor / funcTypeFor / extractHeader
package splitter

import (
	"testing"
)

func TestDocOf(t *testing.T) {
	fn := parseDeclWithComments(t, "package p\n// f doc.\nfunc f() {}")
	if got := docOf(fn); got != "f doc.\n" {
		t.Errorf("docOf func = %q", got)
	}
	gd := parseDeclWithComments(t, "package p\n// T doc.\ntype T int")
	if got := docOf(gd); got != "T doc.\n" {
		t.Errorf("docOf gendecl = %q", got)
	}
}
