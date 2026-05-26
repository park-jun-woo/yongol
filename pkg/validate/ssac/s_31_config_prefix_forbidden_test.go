//ff:func feature=validate type=test control=sequence dimension=3 topic=ssac-structural
//ff:what S-31 — config.* Inputs 금지 검증 (config. 접두사 → ERROR, 다른 접두사 → 통과)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS31ConfigPrefixForbidden(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "post", Line: 5, Inputs: map[string]string{"key": "config.SECRET_KEY"}},
				}},
			},
		}
		diags := s31ConfigPrefixForbidden(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-31]") {
			t.Errorf("Message = %q, want [S-31]", diags[0].Message)
		}
	})
	t.Run("Passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "post", Line: 5, Inputs: map[string]string{"name": "request.Name"}},
				}},
			},
		}
		diags := s31ConfigPrefixForbidden(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
