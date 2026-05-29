//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-67 PASS — bool 반환 predicate Func 통과

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS67EvalFuncSignatureBoolPass(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name: "DoIt",
			Sequences: []parsessac.Sequence{{
				Type:      parsessac.SeqEval,
				Model:     "billing.IsZeroBalance",
				ErrStatus: 402,
			}},
		}},
		ProjectFuncSpecs: []funcspec.FuncSpec{{
			Package:     "billing",
			Name:        "isZeroBalance",
			ReturnTypes: []string{"bool"},
		}},
	}
	if diags := s67EvalFuncSignature(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics for bool predicate, got %d (%v)", len(diags), diags)
	}
}
