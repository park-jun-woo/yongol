//ff:func feature=gen-splitter type=test control=sequence
//ff:what methodReceiver / receiverName / funcIdentifier / genDeclIdentifier / declIdentifier / funcFileName / genDeclFileName / fileNameForDecl
package splitter

import (
	"go/ast"
	"testing"
)

func TestFuncFileName(t *testing.T) {
	plain := declOf(t, "package p\nfunc FindUser() {}").(*ast.FuncDecl)
	if got := funcFileName(plain, ".go"); got != "find_user.go" {
		t.Errorf("plain = %q, want find_user.go", got)
	}
	method := declOf(t, "package p\nfunc (q *Queries) FindUser() {}").(*ast.FuncDecl)
	if got := funcFileName(method, ".sql.go"); got != "queries_find_user.sql.go" {
		t.Errorf("method = %q, want queries_find_user.sql.go", got)
	}
}
