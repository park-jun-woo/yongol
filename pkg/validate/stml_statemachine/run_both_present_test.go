//ff:func feature=validate type=test control=sequence dimension=1 topic=stml-statemachine
//ff:what Run — STML·stateDiagram 둘 다 있고 가드가 합법일 때 진단 없음

package stml_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRunBothPresentClean(t *testing.T) {
	// draft is a valid from-state for ActivateWorkflow → no TM-15, no TM-18.
	fs := &yongol.Fullstack{
		StateDiagrams: []*statemachine.StateDiagram{workflowDiagram()},
		STMLPages: []stml.PageSpec{pageWithAction(stml.ActionBlock{
			OperationID: "ActivateWorkflow",
			EnabledWhen: "workflow.status=draft",
		})},
	}
	if diags := Run(fs); len(diags) != 0 {
		t.Fatalf("expected no diagnostics for legal guard, got %d: %+v", len(diags), diags)
	}
}
