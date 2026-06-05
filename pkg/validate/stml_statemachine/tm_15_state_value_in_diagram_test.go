//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-statemachine
//ff:what TM-15 — 상태값 존재/부재·대소문자 정규화·model 미매칭·빈 가드 분기 검증

package stml_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM15StateValueInDiagram(t *testing.T) {
	stateMap := buildDiagramStateMap([]*statemachine.StateDiagram{workflowDiagram()})

	tests := []struct {
		name      string
		action    stml.ActionBlock
		wantDiags int
	}{
		{
			name:      "value exists in diagram (ok)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow", EnabledWhen: "workflow.status=active"},
			wantDiags: 0,
		},
		{
			name:      "value missing from diagram (error)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow", EnabledWhen: "workflow.status=pending"},
			wantDiags: 1,
		},
		{
			name:      "lowercase model normalizes to PascalCase Symbol (ok)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow", EnabledWhen: "workflow.status=draft"},
			wantDiags: 0,
		},
		{
			name:      "model matches no diagram (no-op)",
			action:    stml.ActionBlock{OperationID: "DoThing", EnabledWhen: "order.status=open"},
			wantDiags: 0,
		},
		{
			name:      "empty enabled-when (skip)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow"},
			wantDiags: 0,
		},
		{
			name:      "compound guard, one bad value (error)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow", EnabledWhen: "workflow.status=draft && workflow.status=bogus"},
			wantDiags: 1,
		},
		{
			name:      "guard parse failure is skipped (TM-17 owns syntax)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow", EnabledWhen: "workflow.status="},
			wantDiags: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := tm15StateValueInDiagram(tt.action, "p.html", stateMap)
			if got := countLevel(diags, "[TM-15]", diagnostic.LevelError); got != tt.wantDiags {
				t.Errorf("TM-15 diags = %d, want %d (%+v)", got, tt.wantDiags, diags)
			}
		})
	}
}
