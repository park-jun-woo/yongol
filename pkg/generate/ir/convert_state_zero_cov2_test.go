//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvert* — direct branch coverage for the per-sequence IR converters
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestConvertState_ZeroCov2(t *testing.T) {
	fs := &yongol.Fullstack{
		StateDiagrams: []*statemachine.StateDiagram{
			{ID: "reservation", Symbol: "Reservation", Transitions: []statemachine.Transition{
				{From: "pending", To: "cancelled", Event: "cancel"},
				{From: "active", To: "cancelled", Event: "cancel"},
			}},
		},
	}
	op := convertState(ssac.Sequence{DiagramID: "reservation", Transition: "cancel"}, fs)
	if op.Kind != OpState || op.State == nil {
		t.Fatalf("expected OpState, got %+v", op)
	}
	if op.State.StatusCode != 409 {
		t.Errorf("default status = %d", op.State.StatusCode)
	}
	if len(op.State.AllowedFromStates) != 2 {
		t.Errorf("allowed from states = %+v", op.State.AllowedFromStates)
	}
	// custom status, nil fs
	op = convertState(ssac.Sequence{DiagramID: "x", Transition: "y", ErrStatus: 423}, nil)
	if op.State.StatusCode != 423 || len(op.State.AllowedFromStates) != 0 {
		t.Errorf("nil-fs state = %+v", op.State)
	}
}
