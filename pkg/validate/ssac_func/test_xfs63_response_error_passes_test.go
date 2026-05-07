//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what XFS-63 PASS — func(req T) (Resp, error) 시그니처는 통과

package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXfs63ResponseErrorPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "DoSomething",
			FileName: "service/do_something.ssac",
			Sequences: []parsessac.Sequence{{
				Type:  "call",
				Model: "billing.DeductCredit",
				Line:  5,
			}},
		}},
		ProjectFuncSpecs: []funcspec.FuncSpec{{
			Package:     "billing",
			Name:        "deductCredit",
			ReturnTypes: []string{"DeductCreditResponse", "error"},
		}},
	}
	diags := xfs63CallFuncSignature(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d (%+v)", len(diags), diags)
	}
}
