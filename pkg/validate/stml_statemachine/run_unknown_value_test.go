//ff:func feature=validate type=test control=sequence dimension=1 topic=stml-statemachine
//ff:what Run — 다이어그램에 없는 상태값을 가드가 참조할 때 TM-15 진단 누적 검증

package stml_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRunBothPresentUnknownValueDiags(t *testing.T) {
	// "pending" is not a state in the diagram at all → TM-15 fires.
	fs := &yongol.Fullstack{
		StateDiagrams: []*statemachine.StateDiagram{workflowDiagram()},
		STMLPages: []stml.PageSpec{pageWithAction(stml.ActionBlock{
			OperationID: "ActivateWorkflow",
			EnabledWhen: "workflow.status=pending",
		})},
	}
	if diags := Run(fs); len(diags) == 0 {
		t.Fatalf("expected a TM-15 diagnostic for unknown state value, got none")
	}
}
