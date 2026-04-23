//ff:func feature=validate type=test control=sequence topic=statemachine-structural
//ff:what ST-1 — diagram 이 이미 로드되어 있으면 규칙 침묵

package statemachine

import (
	"testing"

	smparser "github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestST01ParseSkipsWhenLoaded ensures the rule stays silent when parse already
// succeeded (fs.StateDiagrams non-empty → Run shortcut returns nil).
func TestST01ParseSkipsWhenLoaded(t *testing.T) {
	fs := &yongol.Fullstack{StateDiagrams: []*smparser.StateDiagram{{}}}
	diags := st01Parse(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics when diagrams pre-loaded, got %d", len(diags))
	}
}
