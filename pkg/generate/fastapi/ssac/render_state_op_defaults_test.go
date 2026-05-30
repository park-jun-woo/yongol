//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderStateOpDefaults — renderStateOp 기본값(statusCode 409/메시지/current_state) 분기

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderStateOpDefaults(t *testing.T) {
	// No status input -> statusExpr defaults to "current_state".
	// StatusCode 0 -> 409, empty Message -> default message.
	op := &ir.StateOp{
		Diagram:           "orders",
		Transition:        "Ship",
		AllowedFromStates: []string{"paid"},
	}
	var b strings.Builder
	renderStateOp(&b, op, "")
	got := b.String()
	if !strings.Contains(got, "current_state not in allowed_ship") {
		t.Errorf("expected default current_state expr, got: %s", got)
	}
	if !strings.Contains(got, "status_code=409") {
		t.Errorf("expected default 409, got: %s", got)
	}
	if !strings.Contains(got, `detail="invalid state transition"`) {
		t.Errorf("expected default message, got: %s", got)
	}
}
