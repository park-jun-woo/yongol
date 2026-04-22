//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateStates — StateDiagram ID와 Transition.Event 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// TestPopulateStates_DiagramAndEvents verifies per-diagram event sets are
// keyed by diagram ID.
func TestPopulateStates_DiagramAndEvents(t *testing.T) {
	sd := &statemachine.StateDiagram{
		ID: "order",
		Transitions: []statemachine.Transition{
			{From: "pending", To: "paid", Event: "pay"},
			{From: "paid", To: "shipped", Event: "ship"},
		},
	}
	fs := newMinimalFullstack(withStateDiagrams(sd))
	g := newGround()

	populateStates(g, fs)

	if !g.Lookup["States.diagram"]["order"] {
		t.Errorf("States.diagram missing 'order': %v", g.Lookup["States.diagram"])
	}
	events := g.Lookup["States.event.order"]
	if !events["pay"] || !events["ship"] {
		t.Errorf("States.event.order = %v, want pay+ship", events)
	}
}
