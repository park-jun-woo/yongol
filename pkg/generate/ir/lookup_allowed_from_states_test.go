//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestLookupAllowedFromStates -- diagramID(ID/Symbol) 매칭 및 미스 시 nil 반환 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestLookupAllowedFromStates(t *testing.T) {
	fs := &yongol.Fullstack{
		StateDiagrams: []*statemachine.StateDiagram{
			{
				ID:     "workflow",
				Symbol: "Workflow",
				Transitions: []statemachine.Transition{
					{From: "draft", To: "active", Event: "Activate"},
					{From: "paused", To: "active", Event: "Activate"},
					{From: "active", To: "done", Event: "Finish"},
				},
			},
		},
	}

	tests := []struct {
		name       string
		diagramID  string
		transition string
		want       []string
	}{
		{"match by ID", "workflow", "Activate", []string{"draft", "paused"}},
		{"match by Symbol", "Workflow", "Finish", []string{"active"}},
		{"no diagram match", "missing", "Activate", nil},
		{"diagram match but unknown transition", "workflow", "Nope", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lookupAllowedFromStates(fs, tt.diagramID, tt.transition)
			assertAllowedFromStates(t, got, tt.want)
		})
	}
}
