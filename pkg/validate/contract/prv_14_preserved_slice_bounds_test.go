//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV14PreservedSliceBounds — slice[0] len 가드 누락 오케스트레이션 검증
package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPRV14PreservedSliceBounds(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "h.go")
	writePreserved(t, p,
		"package service\nfunc F(xs []int) int { return xs[0] }\n")
	diags := prv14PreservedSliceBounds([]string{p})
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[PRV-14]") {
		t.Fatalf("expected one PRV-14 diag, got %+v", diags)
	}

	t.Run("guarded is safe", func(t *testing.T) {
		q := filepath.Join(dir, "safe.go")
		writePreserved(t, q,
			"package service\nfunc G(xs []int) int {\n"+
				"  if len(xs) == 0 { return 0 }\n"+
				"  return xs[0]\n}\n")
		if d := prv14PreservedSliceBounds([]string{q}); len(d) != 0 {
			t.Errorf("guarded slice access should be safe, got %+v", d)
		}
	})
}
