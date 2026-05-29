//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV14SliceUnguarded — len 가드 없이 [0] 접근은 ERROR

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV14SliceUnguarded(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nfunc h(items []int) int { return items[0] }\n")
	diags := prv14PreservedSliceBounds([]string{p})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
}
