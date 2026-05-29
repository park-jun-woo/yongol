//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV14SliceGuarded — len 가드 앞선 [0] 접근은 ERROR 없음

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV14SliceGuarded(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nfunc h(items []int) int {\n"+
			"\tif len(items) == 0 { return -1 }\n"+
			"\treturn items[0]\n}\n")
	diags := prv14PreservedSliceBounds([]string{p})
	if len(diags) != 0 {
		t.Fatalf("guarded slice should be safe, got %+v", diags)
	}
}
