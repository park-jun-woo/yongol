//ff:func feature=gen-ir type=test control=sequence
//ff:what TestStateOpAllowedFromStates -- StateOp.AllowedFromStates Mermaid 전이 이식 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
