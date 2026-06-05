//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-statemachine
//ff:what TM-18 — 합법/불법 출발 전이·model 미매칭·빈 가드 분기 검증 (WARNING)

package stml_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM18TransitionValidity(t *testing.T) {
	diagramBySymbol := buildDiagramBySymbol([]*statemachine.StateDiagram{workflowDiagram()})

	tests := []struct {
		name      string
		action    stml.ActionBlock
		wantDiags int
	}{
		{
			name:      "enabled from legal source state (ok)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow", EnabledWhen: "workflow.status=draft"},
			wantDiags: 0,
		},
		{
			name:      "enabled from illegal source state (warning)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow", EnabledWhen: "workflow.status=active"},
			wantDiags: 1,
		},
		{
			name:      "lowercase model normalizes, legal (ok)",
			action:    stml.ActionBlock{OperationID: "ArchiveWorkflow", EnabledWhen: "workflow.status=active"},
			wantDiags: 0,
		},
		{
			name:      "model matches no diagram (no-op)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow", EnabledWhen: "order.status=draft"},
			wantDiags: 0,
		},
		{
			name:      "empty enabled-when (skip)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow"},
			wantDiags: 0,
		},
		{
			name:      "guard parse failure is skipped (TM-17 owns syntax)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow", EnabledWhen: "workflow.status="},
			wantDiags: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := tm18TransitionValidity(tt.action, "p.html", diagramBySymbol)
			if got := countLevel(diags, "[TM-18]", diagnostic.LevelWarning); got != tt.wantDiags {
				t.Errorf("TM-18 diags = %d, want %d (%+v)", got, tt.wantDiags, diags)
			}
		})
	}
}
