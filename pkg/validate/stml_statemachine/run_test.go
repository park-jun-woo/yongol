//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-statemachine
//ff:what Run no-op — STML 또는 stateDiagram 부재 시 진단 없음

package stml_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRunNoOp(t *testing.T) {
	action := stml.ActionBlock{OperationID: "ActivateWorkflow", EnabledWhen: "workflow.status=draft"}

	tests := []struct {
		name string
		fs   *yongol.Fullstack
	}{
		{
			name: "no stml pages",
			fs:   &yongol.Fullstack{StateDiagrams: []*statemachine.StateDiagram{workflowDiagram()}},
		},
		{
			name: "no state diagrams",
			fs:   &yongol.Fullstack{STMLPages: []stml.PageSpec{pageWithAction(action)}},
		},
		{
			name: "both empty",
			fs:   &yongol.Fullstack{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diags := Run(tt.fs); len(diags) != 0 {
				t.Fatalf("expected no diagnostics (no-op), got %d: %+v", len(diags), diags)
			}
		})
	}
}
