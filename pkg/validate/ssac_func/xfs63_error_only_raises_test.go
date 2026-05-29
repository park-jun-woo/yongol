//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what XFS-63 ERROR — func(req T) error 시그니처는 거절

package ssac_func

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXfs63ErrorOnlyRaises(t *testing.T) {
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
			ReturnTypes: []string{"error"},
		}},
	}
	diags := xfs63CallFuncSignature(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XFS-63]") {
		t.Errorf("expected [XFS-63] in message, got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "(error)") {
		t.Errorf("expected actual signature in message, got %q", diags[0].Message)
	}
}
