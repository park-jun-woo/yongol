//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-68 SCOPE — @call 등 다른 시퀀스는 영향 없음

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS68EvalStatusRequiredOnlyAffectsEval(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name: "DoIt",
			Sequences: []parsessac.Sequence{{
				Type:    parsessac.SeqCall,
				Model:   "pkg.DoSomething",
				Message: "x",
			}},
		}},
	}
	if diags := s68EvalStatusRequired(fs); len(diags) != 0 {
		t.Fatalf("S-68 should ignore @call sequences, got %d", len(diags))
	}
}
