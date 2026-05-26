//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-24 — @publish requires Payload 검증 (Inputs/Fields 빈 → ERROR, Inputs 있음 → 통과, 비 publish 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS24PublishPayload(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "publish", Line: 3},
				}},
			},
		}
		diags := s24PublishPayload(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-24]") {
			t.Errorf("Message = %q, want [S-24]", diags[0].Message)
		}
	})
	t.Run("NonPublishSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3},
				}},
			},
		}
		diags := s24PublishPayload(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("PassesWithInputs", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "publish", Line: 3, Inputs: map[string]string{"id": "order.ID"}},
				}},
			},
		}
		diags := s24PublishPayload(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("PassesWithFields", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "publish", Line: 3, Fields: map[string]string{"id": "order.ID"}},
				}},
			},
		}
		diags := s24PublishPayload(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
