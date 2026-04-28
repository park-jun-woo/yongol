//ff:func feature=validate type=test-helper control=sequence topic=hurl-statemachine
//ff:what newWorkflowDiagram — 테스트용 workflow state diagram (draft → active → active)

package hurl_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/statemachine"

// newWorkflowDiagram builds the state diagram: draft -(Activate)-> active -(Execute)-> active.
func newWorkflowDiagram() *statemachine.StateDiagram {
	return &statemachine.StateDiagram{
		ID:           "workflow",
		Symbol:       "Workflow",
		InitialState: "draft",
		States:       []string{"draft", "active"},
		Transitions: []statemachine.Transition{
			{From: "draft", To: "active", Event: "ActivateWorkflow"},
			{From: "active", To: "active", Event: "ExecuteWorkflow"},
		},
	}
}
