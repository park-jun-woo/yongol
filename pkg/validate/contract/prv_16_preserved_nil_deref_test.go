//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestPRV16PreservedNilDeref — Get*/Find* 반환값 즉시 selector 접근 오케스트레이션 검증

package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPRV16PreservedNilDeref(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "h.go")
	writePreserved(t, p,
		"package service\nfunc F(repo R) {\n"+
			"  _ = repo.GetUser().Name\n}\n")
	diags := prv16PreservedNilDeref([]string{p})
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[PRV-16]") {
		t.Fatalf("expected one PRV-16 diag, got %+v", diags)
	}

	t.Run("bound then guarded → safe", func(t *testing.T) {
		q := filepath.Join(dir, "safe.go")
		writePreserved(t, q,
			"package service\nfunc G(repo R) {\n"+
				"  u := repo.GetUser()\n  if u == nil { return }\n  _ = u.Name\n}\n")
		if d := prv16PreservedNilDeref([]string{q}); len(d) != 0 {
			t.Errorf("bound-then-guard should be safe, got %+v", d)
		}
	})
}
