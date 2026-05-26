//ff:func feature=validate type=test control=sequence dimension=3 topic=ssac-structural
//ff:what S-27 — Args.Source 변수 선언 여부 검증 (미선언 → ERROR, 선언됨 → 통과, implicit 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS27VarDeclared(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "put", Line: 5, Args: []ssac.Arg{{Source: "undeclared", Field: "ID"}}},
				}},
			},
		}
		diags := s27VarDeclared(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-27]") {
			t.Errorf("Message = %q, want [S-27]", diags[0].Message)
		}
	})
	t.Run("ImplicitPasses", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "put", Line: 5, Args: []ssac.Arg{{Source: "request", Field: "Name"}}},
				}},
			},
		}
		diags := s27VarDeclared(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 (implicit var should pass)", len(diags))
		}
	})
	t.Run("DeclaredPasses", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID", Result: &ssac.Result{Type: "Order", Var: "order"}},
					{Type: "put", Line: 5, Args: []ssac.Arg{{Source: "order", Field: "ID"}}},
				}},
			},
		}
		diags := s27VarDeclared(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("EmptySourceSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "put", Line: 5, Args: []ssac.Arg{{Source: "", Literal: "42"}}},
				}},
			},
		}
		diags := s27VarDeclared(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 (empty source should be skipped)", len(diags))
		}
	})
}
