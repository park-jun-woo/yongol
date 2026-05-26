//ff:func feature=validate type=test control=sequence dimension=3 topic=ssac-structural
//ff:what S-29 — Inputs 변수 선언 여부 검증 (미선언 → ERROR, implicit 스킵, 빈 Inputs 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS29InputDeclared(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "state", Line: 5, Inputs: map[string]string{"id": "undeclared.ID"}},
				}},
			},
		}
		diags := s29InputDeclared(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-29]") {
			t.Errorf("Message = %q, want [S-29]", diags[0].Message)
		}
	})
	t.Run("EmptyInputsSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "state", Line: 5},
				}},
			},
		}
		diags := s29InputDeclared(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("ImplicitPasses", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "state", Line: 5, Inputs: map[string]string{"name": "request.Name"}},
				}},
			},
		}
		diags := s29InputDeclared(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
