//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-68 PASS — STATUS 명시 시 통과

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS68EvalStatusRequiredPresentPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name: "DoIt",
			Sequences: []parsessac.Sequence{{
				Type:      parsessac.SeqEval,
				Model:     "pkg.IsThing",
				Message:   "Something",
				ErrStatus: 402,
			}},
		}},
	}
	if diags := s68EvalStatusRequired(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d", len(diags))
	}
}
