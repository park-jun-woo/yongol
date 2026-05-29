//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV16NilDerefInline — GetUser().Name 체인은 ERROR

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV16NilDerefInline(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nfunc h(repo interface{ GetUser() *struct{ Name string } }) string { return repo.GetUser().Name }\n")
	diags := prv16PreservedNilDeref([]string{p})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
}
