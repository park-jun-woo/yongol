//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestStatemachineBatch_ZeroCov — ssac_statemachine 빌더/수집기 헬퍼 직접 커버
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestCollectUnmatchedEvents_ZeroCov(t *testing.T) {
	// nil diagram
	if d := collectUnmatchedEvents(nil, map[string]bool{}); d != nil {
		t.Error("nil diagram should return nil")
	}
	diagram := &statemachine.StateDiagram{ID: "wf", File: "states/wf.md", Transitions: []statemachine.Transition{
		{Event: "activate", Line: 2},
		{Event: "activate", Line: 3}, // duplicate event → first-occurrence only
		{Event: "close", Line: 4},
	}}
	funcNames := map[string]bool{"activate": true}
	diags := collectUnmatchedEvents(diagram, funcNames)
	if len(diags) != 1 || diags[0].Line != 4 {
		t.Fatalf("expected 1 diag at line 4, got %v", diags)
	}
	// empty File → synthesized
	d2 := &statemachine.StateDiagram{ID: "wf", Transitions: []statemachine.Transition{{Event: "x", Line: 1}}}
	if d := collectUnmatchedEvents(d2, map[string]bool{}); d[0].File != "states/wf.md" {
		t.Errorf("synth file = %q", d[0].File)
	}
}
