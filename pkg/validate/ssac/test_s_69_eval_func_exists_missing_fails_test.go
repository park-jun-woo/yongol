//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-69 FAIL — Func Spec 누락 시 ERROR

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS69EvalFuncExistsMissingFails(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name: "DoIt",
			Sequences: []parsessac.Sequence{{
				Type:      parsessac.SeqEval,
				Model:     "missing.Predicate",
				ErrStatus: 402,
			}},
		}},
	}
	diags := s69EvalFuncExists(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic for missing func spec, got %d", len(diags))
	}
}
