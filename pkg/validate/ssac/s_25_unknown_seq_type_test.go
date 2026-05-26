//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-25 — unknown sequence type 검증 (미등록 타입 → ERROR, 등록 타입 → 통과)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS25UnknownSeqType(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "foobar", Line: 3},
				}},
			},
		}
		diags := s25UnknownSeqType(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-25]") {
			t.Errorf("Message = %q, want [S-25]", diags[0].Message)
		}
	})
	t.Run("KnownTypePasses", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3},
				}},
			},
		}
		diags := s25UnknownSeqType(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
