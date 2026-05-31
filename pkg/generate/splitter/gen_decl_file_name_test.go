//ff:func feature=gen-splitter type=test control=sequence
//ff:what methodReceiver / receiverName / funcIdentifier / genDeclIdentifier / declIdentifier / funcFileName / genDeclFileName / fileNameForDecl
package splitter

import (
	"go/ast"
	"testing"
)

func TestGenDeclFileName(t *testing.T) {
	typ := declOf(t, "package p\ntype Row struct{}").(*ast.GenDecl)
	if got := genDeclFileName(typ, ".go"); got != "row.go" {
		t.Errorf("type = %q, want row.go", got)
	}
	constBlock := declOf(t, "package p\nconst (\nA=1\nB=2\n)").(*ast.GenDecl)
	if got := genDeclFileName(constBlock, ".go"); got != "consts.go" {
		t.Errorf("const block = %q, want consts.go", got)
	}
	varBlock := declOf(t, "package p\nvar (\nA int\nB int\n)").(*ast.GenDecl)
	if got := genDeclFileName(varBlock, ".go"); got != "vars.go" {
		t.Errorf("var block = %q, want vars.go", got)
	}
}
