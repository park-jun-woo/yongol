//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestScanFileForNilDeref — Get/Find 반환값 즉시 selector 접근 진단 검증

package contract

import (
	"path/filepath"
	"testing"
)

func TestScanFileForNilDeref(t *testing.T) {
	dir := t.TempDir()

	t.Run("GetX().Field flagged", func(t *testing.T) {
		p := filepath.Join(dir, "g.go")
		writePreserved(t, p,
			"package service\nfunc F(repo R) {\n  _ = repo.GetUser().Name\n}\n")
		if d := scanFileForNilDeref(p); len(d) != 1 {
			t.Fatalf("expected 1 diag, got %+v", d)
		}
	})

	t.Run("non Get/Find chain → safe", func(t *testing.T) {
		p := filepath.Join(dir, "ok.go")
		writePreserved(t, p,
			"package service\nfunc G(repo R) {\n  _ = repo.LoadUser().Name\n}\n")
		if d := scanFileForNilDeref(p); len(d) != 0 {
			t.Errorf("expected no diag, got %+v", d)
		}
	})
}
