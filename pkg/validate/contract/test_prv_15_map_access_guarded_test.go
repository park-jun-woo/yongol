//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV15MapAccessGuarded — comma-ok 로 분리된 map 접근은 ERROR 없음

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV15MapAccessGuarded(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\ntype user struct{ Name string }\n\n"+
			"func h(users map[int]*user, id int) string {\n"+
			"\tv, ok := users[id]\n"+
			"\tif !ok || v == nil { return \"\" }\n"+
			"\treturn v.Name\n}\n")
	diags := prv15PreservedMapAccess([]string{p})
	if len(diags) != 0 {
		t.Fatalf("guarded map access should be safe, got %+v", diags)
	}
}
