//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-statemachine
//ff:what buildDiagramStateMap — Symbol→상태집합 변환·다중 다이어그램·중복 상태·빈 입력 검증

package stml_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestBuildDiagramStateMap(t *testing.T) {
	order := &statemachine.StateDiagram{
		ID:     "order",
		Symbol: "Order",
		States: []string{"open", "paid", "open"}, // duplicate collapses in the set
	}
	tests := []struct {
		name     string
		diagrams []*statemachine.StateDiagram
		want     map[string]map[string]bool
	}{
		{
			name:     "nil input yields empty map",
			diagrams: nil,
			want:     map[string]map[string]bool{},
		},
		{
			name:     "single diagram keyed by Symbol",
			diagrams: []*statemachine.StateDiagram{workflowDiagram()},
			want:     map[string]map[string]bool{"Workflow": {"draft": true, "active": true, "archived": true}},
		},
		{
			name:     "duplicate state collapses into single set entry",
			diagrams: []*statemachine.StateDiagram{order},
			want:     map[string]map[string]bool{"Order": {"open": true, "paid": true}},
		},
		{
			name:     "multiple diagrams keyed independently",
			diagrams: []*statemachine.StateDiagram{workflowDiagram(), order},
			want: map[string]map[string]bool{
				"Workflow": {"draft": true, "active": true, "archived": true},
				"Order":    {"open": true, "paid": true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertStateMapEqual(t, buildDiagramStateMap(tt.diagrams), tt.want)
		})
	}
}
