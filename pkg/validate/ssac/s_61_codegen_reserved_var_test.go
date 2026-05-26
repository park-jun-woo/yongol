//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-61 — 결과 변수명이 codegen 예약 식별자와 충돌하면 에러

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS61CodegenReservedVar(t *testing.T) {
	t.Run("Fires_reserved_ctx", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "get", Line: 3, Model: "M.Find",
				Result: &ssac.Result{Var: "ctx", Type: "M"},
			}},
		}}}
		assertDiag(t, s61CodegenReservedVar(fs), "[S-61]")
	})
	t.Run("Fires_reserved_err", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "get", Line: 3, Model: "M.Find",
				Result: &ssac.Result{Var: "err", Type: "M"},
			}},
		}}}
		assertDiag(t, s61CodegenReservedVar(fs), "[S-61]")
	})
	t.Run("Passes_normal_var", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "get", Line: 3, Model: "M.Find",
				Result: &ssac.Result{Var: "course", Type: "Course"},
			}},
		}}}
		assertNoDiag(t, s61CodegenReservedVar(fs), "[S-61]")
	})
	t.Run("Skips_nil_result", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "put", Line: 3, Model: "M.Update",
			}},
		}}}
		assertNoDiag(t, s61CodegenReservedVar(fs), "[S-61]")
	})
	t.Run("Empty_funcs", func(t *testing.T) {
		assertNoDiag(t, s61CodegenReservedVar(&yongol.Fullstack{}), "[S-61]")
	})
}
