//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-1 — @get requires Model 검증 (Model 빈 문자열 → ERROR, Model 있음 → 통과, 비 get 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS01GetModel(t *testing.T) {
	t.Run("missing model fires S-1", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: ""},
				}},
			},
		}
		diags := s01GetModel(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-1]") {
			t.Errorf("Message = %q, want [S-1]", diags[0].Message)
		}
	})
	t.Run("valid model passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID"},
				}},
			},
		}
		diags := s01GetModel(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("non-get skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "post", Line: 5, Model: ""},
				}},
			},
		}
		diags := s01GetModel(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
