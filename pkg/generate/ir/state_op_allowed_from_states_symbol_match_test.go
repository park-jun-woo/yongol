//ff:func feature=gen-ir type=test control=sequence
//ff:what TestStateOpAllowedFromStates -- StateOp.AllowedFromStates Mermaid 전이 이식 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestStateOpAllowedFromStatesSymbolMatch(t *testing.T) {
	fs := &yongol.Fullstack{
		StateDiagrams: []*statemachine.StateDiagram{
			{
				ID:     "workflow",
				Symbol: "Workflow",
				Transitions: []statemachine.Transition{
					{From: "draft", To: "active", Event: "Activate"},
				},
			},
		},
	}

	// DiagramID uses PascalCase Symbol form.
	sf := &ssac.ServiceFunc{
		Name:     "Activate",
		FileName: "activate.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:       ssac.SeqState,
				DiagramID:  "Workflow",
				Inputs:     map[string]string{"Status": "wf.Status"},
				Transition: "Activate",
			},
		},
	}

	plan, err := BuildServicePlan(sf, fs)
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	allowed := plan.Ops[0].State.AllowedFromStates
	if len(allowed) != 1 || allowed[0] != "draft" {
		t.Errorf("AllowedFromStates = %v, want [draft]", allowed)
	}
}
