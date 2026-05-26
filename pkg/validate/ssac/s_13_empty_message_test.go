//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-13 — s13EmptyMessage 검증 (Message 빈 → ERROR, Message 있음 → 통과, 비 empty 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS13EmptyMessage(t *testing.T) {
	t.Run("fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "empty", Line: 3, Message: ""},
				}},
			},
		}
		diags := s13EmptyMessage(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-13]") {
			t.Errorf("Message = %q, want [S-13]", diags[0].Message)
		}
	})
	t.Run("passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "empty", Line: 3, Message: "not found"},
				}},
			},
		}
		diags := s13EmptyMessage(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("non-target skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Message: ""},
				}},
			},
		}
		diags := s13EmptyMessage(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
