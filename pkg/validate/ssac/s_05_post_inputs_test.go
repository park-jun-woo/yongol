//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-5 — @post requires Inputs 검증 (Args/Inputs 둘 다 빈 → ERROR, Args 있음 → 통과)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS05PostInputs(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "post", Line: 3},
				}},
			},
		}
		diags := s05PostInputs(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-5]") {
			t.Errorf("Message = %q, want [S-5]", diags[0].Message)
		}
	})
	t.Run("NonPostSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3},
				}},
			},
		}
		diags := s05PostInputs(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("PassesWithArgs", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "post", Line: 3, Args: []ssac.Arg{{Source: "request", Field: "Name"}}},
				}},
			},
		}
		diags := s05PostInputs(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("PassesWithInputs", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "post", Line: 3, Inputs: map[string]string{"name": "request.Name"}},
				}},
			},
		}
		diags := s05PostInputs(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
