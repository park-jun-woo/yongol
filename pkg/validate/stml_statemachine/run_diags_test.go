//ff:func feature=validate type=test control=sequence dimension=1 topic=stml-statemachine
//ff:what Run — STML·stateDiagram 둘 다 있고 가드가 불법일 때 TM-15/TM-18 진단 누적 검증

package stml_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRunBothPresentDiags(t *testing.T) {
	// "active" is not a value the diagram lacks (it exists), but it is NOT a valid
	// from-state for ActivateWorkflow (only "draft" is), so TM-18 fires.
	fs := &yongol.Fullstack{
		StateDiagrams: []*statemachine.StateDiagram{workflowDiagram()},
		STMLPages: []stml.PageSpec{pageWithAction(stml.ActionBlock{
			OperationID: "ActivateWorkflow",
			EnabledWhen: "workflow.status=active",
		})},
	}
	diags := Run(fs)
	if len(diags) == 0 {
		t.Fatalf("expected at least one diagnostic for illegal transition guard, got none")
	}
}
