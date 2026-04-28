//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-67 FAIL — error 반환 Func 거절 (use @call instead)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS67EvalFuncSignatureErrorReturnFails(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name: "DoIt",
			Sequences: []parsessac.Sequence{{
				Type:      parsessac.SeqEval,
				Model:     "billing.CheckCredits",
				ErrStatus: 402,
			}},
		}},
		ProjectFuncSpecs: []funcspec.FuncSpec{{
			Package:     "billing",
			Name:        "checkCredits",
			ReturnTypes: []string{"error"},
		}},
	}
	diags := s67EvalFuncSignature(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
}
