//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestScanFileForMissingClose — 함수별 리소스 획득/close 매칭 검증
package contract

import (
	"path/filepath"
	"testing"
)

func TestScanFileForMissingClose(t *testing.T) {
	dir := t.TempDir()

	t.Run("missing defer close flagged", func(t *testing.T) {
		p := filepath.Join(dir, "m.go")
		writePreserved(t, p,
			"package service\nfunc F() error {\n"+
				"  f, err := os.Open(\"x\")\n  if err != nil { return err }\n  _ = f\n  return nil\n}\n")
		if d := scanFileForMissingClose(p); len(d) != 1 {
			t.Fatalf("expected 1 diag, got %+v", d)
		}
	})

	t.Run("defer close present → safe", func(t *testing.T) {
		p := filepath.Join(dir, "ok.go")
		writePreserved(t, p,
			"package service\nfunc G() error {\n"+
				"  f, err := os.Open(\"x\")\n  if err != nil { return err }\n  defer f.Close()\n  _ = f\n  return nil\n}\n")
		if d := scanFileForMissingClose(p); len(d) != 0 {
			t.Errorf("expected no diag, got %+v", d)
		}
	})
}
