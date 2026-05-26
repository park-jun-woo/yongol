//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-14 — s14StateDiagramID 검증 (DiagramID 빈 → ERROR, DiagramID 있음 → 통과, 비 state 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS14StateDiagramID(t *testing.T) {
	t.Run("fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "state", Line: 3, DiagramID: ""},
				}},
			},
		}
		diags := s14StateDiagramID(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-14]") {
			t.Errorf("Message = %q, want [S-14]", diags[0].Message)
		}
	})
	t.Run("passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "state", Line: 3, DiagramID: "reservation"},
				}},
			},
		}
		diags := s14StateDiagramID(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("non-target skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, DiagramID: ""},
				}},
			},
		}
		diags := s14StateDiagramID(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
