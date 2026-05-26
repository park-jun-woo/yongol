//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-23 — s23PublishTopic 검증 (Topic 빈 → ERROR, Topic 있음 → 통과, 비 publish 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS23PublishTopic(t *testing.T) {
	t.Run("fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "publish", Line: 3, Topic: ""},
				}},
			},
		}
		diags := s23PublishTopic(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-23]") {
			t.Errorf("Message = %q, want [S-23]", diags[0].Message)
		}
	})
	t.Run("passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "publish", Line: 3, Topic: "order.completed"},
				}},
			},
		}
		diags := s23PublishTopic(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("non-target skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Topic: ""},
				}},
			},
		}
		diags := s23PublishTopic(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
