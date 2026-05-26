//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-67 — @eval 대상 함수 시그니처가 bool 반환이 아니면 에러

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS67EvalFuncSignature(t *testing.T) {
	t.Run("Fires_non_bool_return", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{{
				Name: "X", FileName: "x.ssac", Line: 1,
				Sequences: []ssac.Sequence{{
					Type: "eval", Line: 3, Model: "billing.IsZeroBalance",
				}},
			}},
			ProjectFuncSpecs: []funcspec.FuncSpec{{
				Package: "billing", Name: "isZeroBalance",
				ReturnTypes: []string{"error"},
			}},
		}
		assertDiag(t, s67EvalFuncSignature(fs), "[S-67]")
	})
	t.Run("Fires_no_return", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{{
				Name: "X", FileName: "x.ssac", Line: 1,
				Sequences: []ssac.Sequence{{
					Type: "eval", Line: 3, Model: "billing.IsZeroBalance",
				}},
			}},
			ProjectFuncSpecs: []funcspec.FuncSpec{{
				Package: "billing", Name: "isZeroBalance",
				ReturnTypes: nil,
			}},
		}
		assertDiag(t, s67EvalFuncSignature(fs), "[S-67]")
	})
	t.Run("Passes_bool_return", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{{
				Name: "X", FileName: "x.ssac", Line: 1,
				Sequences: []ssac.Sequence{{
					Type: "eval", Line: 3, Model: "billing.IsZeroBalance",
				}},
			}},
			ProjectFuncSpecs: []funcspec.FuncSpec{{
				Package: "billing", Name: "isZeroBalance",
				ReturnTypes: []string{"bool"},
			}},
		}
		assertNoDiag(t, s67EvalFuncSignature(fs), "[S-67]")
	})
	t.Run("Skips_non_eval", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{{
				Name: "X", FileName: "x.ssac", Line: 1,
				Sequences: []ssac.Sequence{{
					Type: "call", Line: 3, Model: "billing.HoldEscrow",
				}},
			}},
		}
		assertNoDiag(t, s67EvalFuncSignature(fs), "[S-67]")
	})
	t.Run("Skips_missing_spec", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{{
				Name: "X", FileName: "x.ssac", Line: 1,
				Sequences: []ssac.Sequence{{
					Type: "eval", Line: 3, Model: "billing.NoSuchFunc",
				}},
			}},
		}
		assertNoDiag(t, s67EvalFuncSignature(fs), "[S-67]")
	})
	t.Run("Skips_empty_model", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{{
				Name: "X", FileName: "x.ssac", Line: 1,
				Sequences: []ssac.Sequence{{
					Type: "eval", Line: 3, Model: "",
				}},
			}},
		}
		assertNoDiag(t, s67EvalFuncSignature(fs), "[S-67]")
	})
	t.Run("Empty_funcs", func(t *testing.T) {
		assertNoDiag(t, s67EvalFuncSignature(&yongol.Fullstack{}), "[S-67]")
	})
}
