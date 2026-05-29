//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what @eval 시퀀스가 S-25 (unknown sequence type)을 발동하지 않음 — BUG-001 통합 회귀

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestS25EvalDoesNotTriggerUnknownSeqType simulates a parser-emitted @eval
// sequence and asserts s25UnknownSeqType emits no diagnostic. This is the
// integration-level guard for BUG-001 — the bug shipped because no test
// drove a SeqEval value through the validate stage.
func TestS25EvalDoesNotTriggerUnknownSeqType(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "Withdraw",
			FileName: "billing.ssac",
			Sequences: []parsessac.Sequence{{
				Type:      parsessac.SeqEval,
				Model:     "billing.IsZeroBalance",
				ErrStatus: 402,
				Line:      10,
			}},
		}},
	}

	diags := s25UnknownSeqType(fs)
	for _, d := range diags {
		if strings.Contains(d.Message, "[S-25]") {
			t.Fatalf("@eval must not trigger S-25 (BUG-001), got: %s", d.Message)
		}
	}
	if len(diags) != 0 {
		t.Fatalf("expected 0 S-25 diagnostics for @eval, got %d (%v)", len(diags), diags)
	}
}
