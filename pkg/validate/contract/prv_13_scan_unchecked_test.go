//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV13ScanUnchecked — Scan 에러 미체크 ERROR 발생

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV13ScanUnchecked(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nfunc h(row interface{ Scan(...any) error }) int {\n"+
			"\tvar x int\n"+
			"\trow.Scan(&x)\n"+
			"\treturn x\n}\n")
	diags := prv13PreservedScanErr([]string{p})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
}
