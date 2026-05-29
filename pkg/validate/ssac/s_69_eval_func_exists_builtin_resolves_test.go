//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-69 BUILTIN — yongol pkg builtin spec 도 인정

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS69EvalFuncExistsBuiltinResolves(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name: "DoIt",
			Sequences: []parsessac.Sequence{{
				Type:      parsessac.SeqEval,
				Model:     "rate.IsLimited",
				ErrStatus: 429,
			}},
		}},
		YongolPkgSpecs: []funcspec.FuncSpec{{
			Package:     "rate",
			Name:        "isLimited",
			ReturnTypes: []string{"bool"},
		}},
	}
	if diags := s69EvalFuncExists(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics for builtin spec, got %d", len(diags))
	}
}
