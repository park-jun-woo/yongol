//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV16NilDerefBound — 변수 바인딩 후 필드 접근은 ERROR 없음

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV16NilDerefBound(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nfunc h(repo interface{ GetUser() *struct{ Name string } }) string {\n"+
			"\tu := repo.GetUser()\n"+
			"\tif u == nil { return \"\" }\n"+
			"\treturn u.Name\n}\n")
	diags := prv16PreservedNilDeref([]string{p})
	if len(diags) != 0 {
		t.Fatalf("bound + guarded access should be safe, got %+v", diags)
	}
}
