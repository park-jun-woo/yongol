//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestScanFileForCurrentUserAssertion — PRV-11 위반 AssignStmt 수집 검증

package contract

import (
	"path/filepath"
	"testing"
)

func TestScanFileForCurrentUserAssertion(t *testing.T) {
	dir := t.TempDir()

	t.Run("single-value assertion flagged", func(t *testing.T) {
		p := filepath.Join(dir, "c.go")
		writePreserved(t, p,
			"package service\nfunc F(ctx C) {\n  cu := ctx.Value(\"currentUser\").(*model.UserClaim)\n  _ = cu\n}\n")
		if d := scanFileForCurrentUserAssertion(p); len(d) != 1 {
			t.Fatalf("expected 1 diag, got %+v", d)
		}
	})

	t.Run("comma-ok form → safe", func(t *testing.T) {
		p := filepath.Join(dir, "ok.go")
		writePreserved(t, p,
			"package service\nfunc G(ctx C) {\n  cu, ok := ctx.Value(\"currentUser\").(*model.UserClaim)\n  _, _ = cu, ok\n}\n")
		if d := scanFileForCurrentUserAssertion(p); len(d) != 0 {
			t.Errorf("expected no diag, got %+v", d)
		}
	})
}
