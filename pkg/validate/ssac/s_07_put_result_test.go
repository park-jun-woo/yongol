//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-7 — s07PutResult 검증 (Result 있음 → ERROR, nil → 통과, 비 put 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS07PutResult(t *testing.T) {
	t.Run("fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "put", Line: 3, Result: &ssac.Result{Type: "Order", Var: "order"}},
				}},
			},
		}
		diags := s07PutResult(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-7]") {
			t.Errorf("Message = %q, want [S-7]", diags[0].Message)
		}
	})
	t.Run("passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "put", Line: 3, Result: nil},
				}},
			},
		}
		diags := s07PutResult(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("non-target skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Result: &ssac.Result{Type: "Order", Var: "order"}},
				}},
			},
		}
		diags := s07PutResult(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
