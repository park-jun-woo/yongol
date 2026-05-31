//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestAuthStoreCallPrelude_ZeroCov — WrapCalls on/off
package ssac

import (
	"strings"
	"testing"
)

func TestAuthStoreCallPrelude_ZeroCov(t *testing.T) {
	off := &methodGen{WrapCalls: false}
	lines, ctxVar := off.authStoreCallPrelude("auth", "Logout")
	if len(lines) != 0 || ctxVar != "ctx" {
		t.Errorf("WrapCalls off: lines=%v ctx=%q", lines, ctxVar)
	}

	on := &methodGen{WrapCalls: true}
	lines2, ctxVar2 := on.authStoreCallPrelude("auth", "Logout")
	if ctxVar2 != "callCtx" {
		t.Errorf("WrapCalls on: ctx=%q want callCtx", ctxVar2)
	}
	if len(lines2) != 1 || !strings.Contains(lines2[0], `otel.Tracer("ssac").Start(ctx, "call.auth.Logout")`) {
		t.Errorf("WrapCalls on: lines=%v", lines2)
	}
}
