//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestScanFileForMapAccess — 가드 없는 map[key].Sel 패턴 진단 수집 검증

package contract

import (
	"path/filepath"
	"testing"
)

func TestScanFileForMapAccess(t *testing.T) {
	dir := t.TempDir()

	t.Run("inline index-select flagged", func(t *testing.T) {
		p := filepath.Join(dir, "m.go")
		writePreserved(t, p,
			"package service\nfunc F(m map[string]T, k string) {\n  _ = m[k].Field\n}\n")
		if d := scanFileForMapAccess(p); len(d) != 1 {
			t.Fatalf("expected 1 diag, got %+v", d)
		}
	})

	t.Run("plain selector → safe", func(t *testing.T) {
		p := filepath.Join(dir, "ok.go")
		writePreserved(t, p,
			"package service\nfunc G(v T) {\n  _ = v.Field\n}\n")
		if d := scanFileForMapAccess(p); len(d) != 0 {
			t.Errorf("expected no diag, got %+v", d)
		}
	})
}
