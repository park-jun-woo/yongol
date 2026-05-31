//ff:func feature=gen-splitter type=test control=sequence
//ff:what methodReceiver / receiverName / funcIdentifier / genDeclIdentifier / declIdentifier / funcFileName / genDeclFileName / fileNameForDecl
package splitter

import (
	"testing"
)

func TestFileNameForDecl(t *testing.T) {
	fn := declOf(t, "package p\nfunc FindUser() {}")
	if got := fileNameForDecl(fn, ToolSQLC, false); got != "find_user.sql.go" {
		t.Errorf("func = %q, want find_user.sql.go", got)
	}
	typ := declOf(t, "package p\ntype Row struct{}")
	if got := fileNameForDecl(typ, ToolSQLC, true); got != "row.model.go" {
		t.Errorf("models type = %q, want row.model.go", got)
	}
	if got := fileNameForDecl(fn, ToolOAPICodegen, false); got != "find_user.gen.go" {
		t.Errorf("oapi func = %q, want find_user.gen.go", got)
	}
}
