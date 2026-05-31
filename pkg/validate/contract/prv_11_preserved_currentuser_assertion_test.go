//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV11PreservedCurrentUserAssertion — currentUser 단일대입 단언 오케스트레이션 검증
package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPRV11PreservedCurrentUserAssertion(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "h.go")
	writePreserved(t, p,
		"package service\nfunc F(ctx C) {\n"+
			"  cu := ctx.Value(\"currentUser\").(*model.UserClaim)\n"+
			"  _ = cu\n}\n")
	diags := prv11PreservedCurrentUserAssertion([]string{p})
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[PRV-11]") {
		t.Fatalf("expected one PRV-11 diag, got %+v", diags)
	}

	t.Run("comma-ok form is safe", func(t *testing.T) {
		q := filepath.Join(dir, "safe.go")
		writePreserved(t, q,
			"package service\nfunc G(ctx C) {\n"+
				"  cu, ok := ctx.Value(\"currentUser\").(*model.UserClaim)\n"+
				"  _, _ = cu, ok\n}\n")
		if d := prv11PreservedCurrentUserAssertion([]string{q}); len(d) != 0 {
			t.Errorf("comma-ok form should be safe, got %+v", d)
		}
	})
}
