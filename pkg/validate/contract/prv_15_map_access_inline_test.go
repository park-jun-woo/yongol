//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV15MapAccessInline — `m[k].Field` inline selector 접근은 ERROR

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV15MapAccessInline(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\ntype user struct{ Name string }\n\n"+
			"func h(users map[int]*user, id int) string { return users[id].Name }\n")
	diags := prv15PreservedMapAccess([]string{p})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
}
