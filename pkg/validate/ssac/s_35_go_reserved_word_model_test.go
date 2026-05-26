//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-35 — Go 예약어를 result 타입명으로 사용 금지 검증 (예약어 → ERROR, 비예약어 → 통과, nil Result 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS35GoReservedWordModel(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Result: &ssac.Result{Type: "func", Var: "f"}},
				}},
			},
		}
		diags := s35GoReservedWordModel(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-35]") {
			t.Errorf("Message = %q, want [S-35]", diags[0].Message)
		}
	})
	t.Run("NilResultSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Result: nil},
				}},
			},
		}
		diags := s35GoReservedWordModel(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("Passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Result: &ssac.Result{Type: "Order", Var: "order"}},
				}},
			},
		}
		diags := s35GoReservedWordModel(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
