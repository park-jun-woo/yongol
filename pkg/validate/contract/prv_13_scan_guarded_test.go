//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV13ScanGuarded — 인라인 err 체크된 Scan 은 ERROR 없음

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV13ScanGuarded(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nfunc h(row interface{ Scan(...any) error }) error {\n"+
			"\tvar x int\n"+
			"\tif err := row.Scan(&x); err != nil { return err }\n"+
			"\t_ = x\n"+
			"\treturn nil\n}\n")
	diags := prv13PreservedScanErr([]string{p})
	if len(diags) != 0 {
		t.Fatalf("guarded Scan should be safe, got %+v", diags)
	}
}
