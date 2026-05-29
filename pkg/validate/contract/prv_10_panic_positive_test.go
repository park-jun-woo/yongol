//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV10PanicPositive — preserved 파일 body 에 panic 발견 시 ERROR 발생

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV10PanicPositive(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nfunc ActivateWorkflow() { panic(\"boom\") }\n")
	diags := prv10PreservedPanic([]string{p})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
}
