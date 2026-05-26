//ff:func feature=validate type=test control=sequence dimension=3 topic=ssac-structural
//ff:what S-30 — @response Fields 변수 선언 여부 검증 (미선언 → ERROR, implicit 스킵, 빈 Fields 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS30MessageDeclared(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "response", Line: 5, Fields: map[string]string{"name": "undeclared.Name"}},
				}},
			},
		}
		diags := s30MessageDeclared(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-30]") {
			t.Errorf("Message = %q, want [S-30]", diags[0].Message)
		}
	})
	t.Run("EmptyFieldsSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "response", Line: 5},
				}},
			},
		}
		diags := s30MessageDeclared(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("ImplicitPasses", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "response", Line: 5, Fields: map[string]string{"name": "request.Name"}},
				}},
			},
		}
		diags := s30MessageDeclared(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
