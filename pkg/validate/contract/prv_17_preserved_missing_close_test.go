//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV17PreservedMissingClose — 리소스 획득 후 defer Close 누락 오케스트레이션 검증
package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPRV17PreservedMissingClose(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "h.go")
	writePreserved(t, p,
		"package service\nfunc F() error {\n"+
			"  f, err := os.Open(\"x\")\n"+
			"  if err != nil { return err }\n"+
			"  _ = f\n  return nil\n}\n")
	diags := prv17PreservedMissingClose([]string{p})
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[PRV-17]") {
		t.Fatalf("expected one PRV-17 diag, got %+v", diags)
	}

	t.Run("defer close → safe", func(t *testing.T) {
		q := filepath.Join(dir, "safe.go")
		writePreserved(t, q,
			"package service\nfunc G() error {\n"+
				"  f, err := os.Open(\"x\")\n"+
				"  if err != nil { return err }\n"+
				"  defer f.Close()\n  _ = f\n  return nil\n}\n")
		if d := prv17PreservedMissingClose([]string{q}); len(d) != 0 {
			t.Errorf("defer close should be safe, got %+v", d)
		}
	})
}
