//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-68 FAIL — STATUS 누락 시 ERROR

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS68EvalStatusRequiredMissingFails(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name: "DoIt",
			Sequences: []parsessac.Sequence{{
				Type:    parsessac.SeqEval,
				Model:   "pkg.IsThing",
				Message: "Something",
			}},
		}},
	}
	diags := s68EvalStatusRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic for missing STATUS, got %d", len(diags))
	}
}
