//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-67 SILENT — Func Spec 누락 시 침묵 (S-69 가 보고)

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS67EvalFuncSignatureSilentWhenSpecMissing(t *testing.T) {
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
	if diags := s67EvalFuncSignature(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics (S-69 reports missing spec), got %d", len(diags))
	}
}
