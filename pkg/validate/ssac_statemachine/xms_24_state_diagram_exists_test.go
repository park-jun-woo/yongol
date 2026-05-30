//ff:func feature=validate type=test control=selection topic=states
//ff:what TestXms24StateDiagramExists — XMS-24 @state 참조 diagram 존재 검증

package ssac_statemachine

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXms24StateDiagramExists(t *testing.T) {
	t.Run("missing diagram fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{{
				Name:      "CancelOrder",
				FileName:  "service/order/cancel.ssac",
				Sequences: []ssac.Sequence{{Type: "state", DiagramID: "order", Line: 4}},
			}},
		}
		diags := xms24StateDiagramExists(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[XMS-24]") {
			t.Errorf("unexpected message: %q", diags[0].Message)
		}
	})

	t.Run("existing diagram passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			StateDiagrams: []*statemachine.StateDiagram{{ID: "order"}},
			ServiceFuncs: []ssac.ServiceFunc{{
				Name:      "CancelOrder",
				Sequences: []ssac.Sequence{{Type: "state", DiagramID: "order"}},
			}},
		}
		if diags := xms24StateDiagramExists(fs); len(diags) != 0 {
			t.Errorf("expected no diagnostics, got %d", len(diags))
		}
	})
}
