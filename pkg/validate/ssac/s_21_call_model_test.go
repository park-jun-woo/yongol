//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-21 — s21CallModel 검증 (Model 빈 → ERROR, Model 있음 → 통과, 비 call 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS21CallModel(t *testing.T) {
	t.Run("fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "call", Line: 3, Model: ""},
				}},
			},
		}
		diags := s21CallModel(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-21]") {
			t.Errorf("Message = %q, want [S-21]", diags[0].Message)
		}
	})
	t.Run("passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "call", Line: 3, Model: "auth.VerifyPassword"},
				}},
			},
		}
		diags := s21CallModel(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("non-target skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: ""},
				}},
			},
		}
		diags := s21CallModel(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
