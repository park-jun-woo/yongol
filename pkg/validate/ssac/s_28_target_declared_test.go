//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-28 — Target 변수 선언 여부 검증 (미선언 → ERROR, 선언됨 → 통과, implicit 스킵, 빈 Target 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS28TargetDeclared(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "empty", Line: 5, Target: "undeclared.ID"},
				}},
			},
		}
		diags := s28TargetDeclared(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-28]") {
			t.Errorf("Message = %q, want [S-28]", diags[0].Message)
		}
	})
	t.Run("EmptyTargetSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "empty", Line: 5, Target: ""},
				}},
			},
		}
		diags := s28TargetDeclared(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("ImplicitPasses", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "empty", Line: 5, Target: "request.ID"},
				}},
			},
		}
		diags := s28TargetDeclared(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 (implicit var should pass)", len(diags))
		}
	})
	t.Run("DeclaredPasses", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Result: &ssac.Result{Type: "Order", Var: "order"}},
					{Type: "empty", Line: 5, Target: "order"},
				}},
			},
		}
		diags := s28TargetDeclared(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
