//ff:func feature=validate type=test control=sequence dimension=1 topic=stml-statemachine
//ff:what workflowDiagram — 테스트용 workflow stateDiagram(draft→active→archived) 생성

package stml_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/statemachine"

func workflowDiagram() *statemachine.StateDiagram {
	return &statemachine.StateDiagram{
		ID:     "workflow",
		Symbol: "Workflow",
		States: []string{"draft", "active", "archived"},
		Transitions: []statemachine.Transition{
			{From: "draft", To: "active", Event: "ActivateWorkflow"},
			{From: "active", To: "archived", Event: "ArchiveWorkflow"},
		},
	}
}
