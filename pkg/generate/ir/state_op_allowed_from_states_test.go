//ff:func feature=gen-ir type=test control=sequence
//ff:what TestStateOpAllowedFromStates -- StateOp.AllowedFromStates Mermaid 전이 이식 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestStateOpAllowedFromStates(t *testing.T) {
	fs := &yongol.Fullstack{
		StateDiagrams: []*statemachine.StateDiagram{
			{
				ID:     "workflow",
				Symbol: "Workflow",
				Transitions: []statemachine.Transition{
					{From: "[*]", To: "draft", Event: "create"},
					{From: "draft", To: "active", Event: "ActivateWorkflow"},
					{From: "active", To: "archived", Event: "ArchiveWorkflow"},
					{From: "draft", To: "archived", Event: "ArchiveWorkflow"},
				},
			},
		},
	}

	sf := &ssac.ServiceFunc{
		Name:     "ArchiveWorkflow",
		FileName: "archive_workflow.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:       ssac.SeqState,
				DiagramID:  "workflow",
				Inputs:     map[string]string{"Status": "wf.Status"},
				Transition: "ArchiveWorkflow",
				ErrStatus:  409,
			},
		},
	}

	plan, err := BuildServicePlan(sf, fs)
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	stateOp := plan.Ops[0]
	if stateOp.Kind != OpState {
		t.Fatalf("Ops[0].Kind = %d, want OpState", stateOp.Kind)
	}
	allowed := stateOp.State.AllowedFromStates
	if len(allowed) != 2 {
		t.Fatalf("len(AllowedFromStates) = %d, want 2", len(allowed))
	}
	// ValidFromStates returns them in transition order.
	found := map[string]bool{}
	for _, s := range allowed {
		found[s] = true
	}
	if !found["active"] || !found["draft"] {
		t.Errorf("AllowedFromStates = %v, want [active, draft]", allowed)
	}
}

func TestStateOpAllowedFromStatesNoDiagram(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "ArchiveWorkflow",
		FileName: "archive_workflow.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:       ssac.SeqState,
				DiagramID:  "workflow",
				Inputs:     map[string]string{"Status": "wf.Status"},
				Transition: "ArchiveWorkflow",
				ErrStatus:  409,
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	stateOp := plan.Ops[0]
	if stateOp.State.AllowedFromStates != nil {
		t.Errorf("AllowedFromStates = %v, want nil when no diagrams", stateOp.State.AllowedFromStates)
	}
}

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
