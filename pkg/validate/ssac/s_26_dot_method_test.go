//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-26 — Model.Method 형식 검증 (도트 없음 → ERROR, 정상 → 통과, 비 CRUD 스킵, 빈 Model 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS26DotMethod(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "NoMethod"},
				}},
			},
		}
		diags := s26DotMethod(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-26]") {
			t.Errorf("Message = %q, want [S-26]", diags[0].Message)
		}
	})
	t.Run("NonCRUDSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "auth", Line: 3, Model: "NoMethod"},
				}},
			},
		}
		diags := s26DotMethod(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 (non-CRUD should be skipped)", len(diags))
		}
	})
	t.Run("EmptyModelSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: ""},
				}},
			},
		}
		diags := s26DotMethod(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 (empty model deferred to S-1)", len(diags))
		}
	})
	t.Run("Passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID"},
				}},
			},
		}
		diags := s26DotMethod(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
