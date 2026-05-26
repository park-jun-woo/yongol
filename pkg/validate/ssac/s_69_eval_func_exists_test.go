//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-69 — @eval 대상 함수가 FuncSpec 에 없으면 에러

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS69EvalFuncExists(t *testing.T) {
	t.Run("Fires_unknown_func", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{{
				Name: "X", FileName: "x.ssac", Line: 1,
				Sequences: []ssac.Sequence{{
					Type: "eval", Line: 3, Model: "billing.NoSuchFunc",
				}},
			}},
		}
		assertDiag(t, s69EvalFuncExists(fs), "[S-69]")
	})
	t.Run("Passes_known_project_func", func(t *testing.T) {
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
		assertNoDiag(t, s69EvalFuncExists(fs), "[S-69]")
	})
	t.Run("Passes_known_builtin_func", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{{
				Name: "X", FileName: "x.ssac", Line: 1,
				Sequences: []ssac.Sequence{{
					Type: "eval", Line: 3, Model: "text.IsEmpty",
				}},
			}},
			YongolPkgSpecs: []funcspec.FuncSpec{{
				Package: "text", Name: "isEmpty",
				ReturnTypes: []string{"bool"},
			}},
		}
		assertNoDiag(t, s69EvalFuncExists(fs), "[S-69]")
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
		assertNoDiag(t, s69EvalFuncExists(fs), "[S-69]")
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
		assertNoDiag(t, s69EvalFuncExists(fs), "[S-69]")
	})
	t.Run("Empty_funcs", func(t *testing.T) {
		assertNoDiag(t, s69EvalFuncExists(&yongol.Fullstack{}), "[S-69]")
	})
}
