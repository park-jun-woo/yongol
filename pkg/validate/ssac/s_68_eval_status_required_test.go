//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-68 — @eval 에 STATUS 코드가 명시되지 않으면 에러

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS68EvalStatusRequired(t *testing.T) {
	t.Run("Fires_missing_status", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "eval", Line: 3, Model: "billing.IsZeroBalance", ErrStatus: 0,
			}},
		}}}
		assertDiag(t, s68EvalStatusRequired(fs), "[S-68]")
	})
	t.Run("Passes_with_status", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "eval", Line: 3, Model: "billing.IsZeroBalance", ErrStatus: 402,
			}},
		}}}
		assertNoDiag(t, s68EvalStatusRequired(fs), "[S-68]")
	})
	t.Run("Skips_non_eval", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "call", Line: 3, ErrStatus: 0,
			}},
		}}}
		assertNoDiag(t, s68EvalStatusRequired(fs), "[S-68]")
	})
	t.Run("Empty_funcs", func(t *testing.T) {
		assertNoDiag(t, s68EvalStatusRequired(&yongol.Fullstack{}), "[S-68]")
	})
}
