//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-62 — 선언된 결과 변수가 후속 시퀀스에서 참조되지 않으면 에러

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS62UnusedResultVar(t *testing.T) {
	t.Run("Fires_unused_var", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{
				{Type: "get", Line: 3, Model: "M.Find", Result: &ssac.Result{Var: "item", Type: "M"}},
				{Type: "response", Line: 5, Fields: map[string]string{"other": "other"}},
			},
		}}}
		assertDiag(t, s62UnusedResultVar(fs), "[S-62]")
	})
	t.Run("Passes_used_in_inputs", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{
				{Type: "get", Line: 3, Model: "M.Find", Result: &ssac.Result{Var: "item", Type: "M"}},
				{Type: "put", Line: 5, Inputs: map[string]string{"id": "item.ID"}},
			},
		}}}
		assertNoDiag(t, s62UnusedResultVar(fs), "[S-62]")
	})
	t.Run("Passes_used_in_fields", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{
				{Type: "get", Line: 3, Model: "M.Find", Result: &ssac.Result{Var: "item", Type: "M"}},
				{Type: "response", Line: 5, Fields: map[string]string{"item": "item"}},
			},
		}}}
		assertNoDiag(t, s62UnusedResultVar(fs), "[S-62]")
	})
	t.Run("Passes_used_in_target", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{
				{Type: "get", Line: 3, Model: "M.Find", Result: &ssac.Result{Var: "item", Type: "M"}},
				{Type: "empty", Line: 5, Target: "item"},
			},
		}}}
		assertNoDiag(t, s62UnusedResultVar(fs), "[S-62]")
	})
	t.Run("Skips_nil_result", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{
				{Type: "put", Line: 3, Model: "M.Update"},
			},
		}}}
		assertNoDiag(t, s62UnusedResultVar(fs), "[S-62]")
	})
	t.Run("Skips_underscore", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{
				{Type: "get", Line: 3, Model: "M.Find", Result: &ssac.Result{Var: "_", Type: "M"}},
			},
		}}}
		assertNoDiag(t, s62UnusedResultVar(fs), "[S-62]")
	})
	t.Run("Empty_funcs", func(t *testing.T) {
		assertNoDiag(t, s62UnusedResultVar(&yongol.Fullstack{}), "[S-62]")
	})
}
