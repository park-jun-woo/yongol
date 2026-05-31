//ff:func feature=gen-splitter type=test control=sequence
//ff:what docOf / funcDoc / genDeclDoc / detectControl / controlFor / funcTypeFor / extractHeader
package splitter

import (
	"testing"
)

func TestControlFor(t *testing.T) {
	fn := parseDeclWithComments(t, "package p\nfunc f() { for {} }")
	if c, d := controlFor(fn); c != "iteration" || d != "1" {
		t.Errorf("func loop = (%q,%q)", c, d)
	}
	gd := parseDeclWithComments(t, "package p\ntype T int")
	if c, d := controlFor(gd); c != "sequence" || d != "" {
		t.Errorf("gendecl = (%q,%q)", c, d)
	}
}
