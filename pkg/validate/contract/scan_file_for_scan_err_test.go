//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestScanFileForScanErr — sql.Row/Rows.Scan 에러 누락 진단 수집 검증

package contract

import (
	"path/filepath"
	"testing"
)

func TestScanFileForScanErr(t *testing.T) {
	dir := t.TempDir()

	t.Run("ignored scan error flagged", func(t *testing.T) {
		p := filepath.Join(dir, "s.go")
		writePreserved(t, p,
			"package service\nfunc F(row R) {\n  var x int\n  row.Scan(&x)\n  _ = x\n}\n")
		if d := scanFileForScanErr(p); len(d) != 1 {
			t.Fatalf("expected 1 diag, got %+v", d)
		}
	})

	t.Run("guarded scan → safe", func(t *testing.T) {
		p := filepath.Join(dir, "ok.go")
		writePreserved(t, p,
			"package service\nfunc G(row R) error {\n  var x int\n  if err := row.Scan(&x); err != nil { return err }\n  _ = x\n  return nil\n}\n")
		if d := scanFileForScanErr(p); len(d) != 0 {
			t.Errorf("expected no diag, got %+v", d)
		}
	})
}
