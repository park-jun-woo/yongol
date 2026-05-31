//ff:func feature=gen-splitter type=test control=sequence
//ff:what docOf / funcDoc / genDeclDoc / detectControl / controlFor / funcTypeFor / extractHeader
package splitter

import (
	"go/ast"
	"testing"
)

func TestGenDeclDoc(t *testing.T) {
	wrapped := parseDeclWithComments(t, "package p\n// Row is a row.\ntype Row struct{}").(*ast.GenDecl)
	if got := genDeclDoc(wrapped); got != "Row is a row.\n" {
		t.Errorf("wrapped doc = %q", got)
	}
	// doc on inner TypeSpec
	inner := parseDeclWithComments(t, "package p\ntype (\n// Inner doc.\nInner struct{}\n)").(*ast.GenDecl)
	if got := genDeclDoc(inner); got != "Inner doc.\n" {
		t.Errorf("inner spec doc = %q", got)
	}
	multi := parseDeclWithComments(t, "package p\ntype (\nA struct{}\nB struct{}\n)").(*ast.GenDecl)
	if got := genDeclDoc(multi); got != "" {
		t.Errorf("multi-spec doc = %q, want empty", got)
	}
}
