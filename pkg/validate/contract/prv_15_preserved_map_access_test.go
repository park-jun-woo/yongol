//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV15PreservedMapAccess — 가드 없는 map[k].Sel 오케스트레이션 검증
package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPRV15PreservedMapAccess(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "h.go")
	writePreserved(t, p,
		"package service\nfunc F(m map[string]T, k string) {\n"+
			"  _ = m[k].Field\n}\n")
	diags := prv15PreservedMapAccess([]string{p})
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[PRV-15]") {
		t.Fatalf("expected one PRV-15 diag, got %+v", diags)
	}

	t.Run("no inline index-select → safe", func(t *testing.T) {
		q := filepath.Join(dir, "safe.go")
		writePreserved(t, q,
			"package service\nfunc G(m map[string]T, k string) {\n"+
				"  v, ok := m[k]\n  if !ok { return }\n  _ = v.Field\n}\n")
		if d := prv15PreservedMapAccess([]string{q}); len(d) != 0 {
			t.Errorf("split access should be safe, got %+v", d)
		}
	})
}
