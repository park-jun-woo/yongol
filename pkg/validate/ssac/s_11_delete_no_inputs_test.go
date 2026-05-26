//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-11 — @delete with no inputs WARNING 검증 (Inputs 빈 → WARNING, Args 있음 → 통과)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS11DeleteNoInputs(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "delete", Line: 3},
				}},
			},
		}
		diags := s11DeleteNoInputs(fs)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-11]") {
			t.Errorf("Message = %q, want [S-11]", diags[0].Message)
		}
		if diags[0].Level != diagnostic.LevelWarning {
			t.Errorf("Level = %q, want WARNING", diags[0].Level)
		}
	})
	t.Run("NonDeleteSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3},
				}},
			},
		}
		diags := s11DeleteNoInputs(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("SuppressWarn", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "delete", Line: 3, SuppressWarn: true},
				}},
			},
		}
		diags := s11DeleteNoInputs(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 (SuppressWarn should skip)", len(diags))
		}
	})
	t.Run("Passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "delete", Line: 3, Args: []ssac.Arg{{Source: "request", Field: "ID"}}},
				}},
			},
		}
		diags := s11DeleteNoInputs(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
