//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-69 PASS — Func Spec 가 존재하면 통과

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS69EvalFuncExistsPasses(t *testing.T) {
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
	if diags := s69EvalFuncExists(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d", len(diags))
	}
}
