//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV13PreservedScanErr — sql.Scan 에러 무시 오케스트레이션 검증
package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPRV13PreservedScanErr(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "h.go")
	writePreserved(t, p,
		"package service\nfunc F(row R) {\n"+
			"  var x int\n"+
			"  row.Scan(&x)\n"+
			"  _ = x\n}\n")
	diags := prv13PreservedScanErr([]string{p})
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[PRV-13]") {
		t.Fatalf("expected one PRV-13 diag, got %+v", diags)
	}

	t.Run("guarded is safe", func(t *testing.T) {
		q := filepath.Join(dir, "safe.go")
		writePreserved(t, q,
			"package service\nfunc G(row R) error {\n"+
				"  var x int\n"+
				"  if err := row.Scan(&x); err != nil { return err }\n"+
				"  _ = x\n  return nil\n}\n")
		if d := prv13PreservedScanErr([]string{q}); len(d) != 0 {
			t.Errorf("guarded scan should be safe, got %+v", d)
		}
	})
}
