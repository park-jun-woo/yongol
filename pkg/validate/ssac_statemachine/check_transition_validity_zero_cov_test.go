//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestStatemachineBatch_ZeroCov — ssac_statemachine 빌더/수집기 헬퍼 직접 커버
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestCheckTransitionValidity_ZeroCov(t *testing.T) {
	diagram := &statemachine.StateDiagram{ID: "wf", Transitions: []statemachine.Transition{
		{From: "draft", To: "active", Event: "activate", Line: 2},
	}}
	byID := map[string]*statemachine.StateDiagram{"wf": diagram}
	fn := ssac.ServiceFunc{Name: "Fn", FileName: "f.ssac"}

	// non-state seq
	if _, ok := checkTransitionValidity(fn, ssac.Sequence{Type: "get"}, byID); ok {
		t.Error("non-state should be skipped")
	}
	// unknown diagram
	if _, ok := checkTransitionValidity(fn, ssac.Sequence{Type: "state", DiagramID: "nope"}, byID); ok {
		t.Error("unknown diagram should be skipped")
	}
	// valid transition
	if _, ok := checkTransitionValidity(fn, ssac.Sequence{Type: "state", DiagramID: "wf", Transition: "activate"}, byID); ok {
		t.Error("valid transition should not raise")
	}
	// invalid transition → diag
	d, ok := checkTransitionValidity(fn, ssac.Sequence{Type: "state", DiagramID: "wf", Transition: "bogus", Line: 9}, byID)
	if !ok {
		t.Fatal("invalid transition should raise")
	}
	if d.Line != 9 || d.OperationID != "Fn" {
		t.Errorf("diag = %+v", d)
	}
	// file synthesized when empty
	d2, ok := checkTransitionValidity(ssac.ServiceFunc{Name: "Fn"}, ssac.Sequence{Type: "state", DiagramID: "wf", Transition: "bogus"}, byID)
	if !ok || d2.File != "ssac/Fn.ssac" {
		t.Errorf("synth file = %q", d2.File)
	}
}
