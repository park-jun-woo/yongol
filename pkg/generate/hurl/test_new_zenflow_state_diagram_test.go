//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what newZenflowStateDiagram — ZenFlow workflow.md 미러링 state machine 생성

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// newZenflowStateDiagram returns the ZenFlow workflow state machine used
// across Phase004 tests. Mirrors examples/zenflow/specs/states/workflow.md.
func newZenflowStateDiagram() *statemachine.StateDiagram {
	return &statemachine.StateDiagram{
		ID:           "workflow",
		Symbol:       "Workflow",
		InitialState: "draft",
		States:       []string{"draft", "active", "paused", "archived"},
		Transitions: []statemachine.Transition{
			{From: "draft", To: "active", Event: "ActivateWorkflow"},
			{From: "paused", To: "active", Event: "ActivateWorkflow"},
			{From: "active", To: "paused", Event: "PauseWorkflow"},
			{From: "active", To: "archived", Event: "ArchiveWorkflow"},
			{From: "active", To: "active", Event: "ExecuteWorkflow"},
		},
	}
}
