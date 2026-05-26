//ff:func feature=validate type=test control=sequence dimension=3 topic=ssac-structural
//ff:what S-32 — @publish Inputs 에 query 금지 검증 (query 사용 → ERROR, 비 publish 스킵, 다른 값 통과)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS32PublishForbidden(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "publish", Line: 5, Inputs: map[string]string{"q": "query"}},
				}},
			},
		}
		diags := s32PublishForbidden(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-32]") {
			t.Errorf("Message = %q, want [S-32]", diags[0].Message)
		}
	})
	t.Run("NonPublishSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "post", Line: 5, Inputs: map[string]string{"q": "query"}},
				}},
			},
		}
		diags := s32PublishForbidden(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("Passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "publish", Line: 5, Inputs: map[string]string{"id": "order.ID"}},
				}},
			},
		}
		diags := s32PublishForbidden(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
