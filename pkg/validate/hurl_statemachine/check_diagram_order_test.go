//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-statemachine
//ff:what checkDiagramOrder — state diagram에 대해 hurl entries의 전이 순서 검증

package hurl_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestCheckDiagramOrder(t *testing.T) {
	d := &statemachine.StateDiagram{
		ID:           "order",
		InitialState: "draft",
		Transitions: []statemachine.Transition{
			{From: "draft", To: "submitted", Event: "submitOrder"},
			{From: "submitted", To: "approved", Event: "approveOrder"},
		},
	}
	opID := map[string]string{
		"POST /orders":                "submitOrder",
		"POST /orders/:param/approve": "approveOrder",
	}

	cases := []struct {
		name      string
		diagram   *statemachine.StateDiagram
		entries   []hurl.HurlEntry
		wantCount int
	}{
		{name: "nil_diagram", diagram: nil, wantCount: 0},
		{name: "empty_initial_state", diagram: &statemachine.StateDiagram{}, wantCount: 0},
		{
			name:    "correct_order_no_diag",
			diagram: d,
			entries: []hurl.HurlEntry{
				{Method: "POST", Path: "/orders", File: "t.hurl", Line: 1},
				{Method: "POST", Path: "/orders/1/approve", File: "t.hurl", Line: 5},
			},
			wantCount: 0,
		},
		{
			name:    "wrong_order_produces_warning",
			diagram: d,
			entries: []hurl.HurlEntry{
				{Method: "POST", Path: "/orders/1/approve", File: "t.hurl", Line: 1},
			},
			wantCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runDiagCodeCase(t, checkDiagramOrder(c.entries, opID, c.diagram), c.wantCount, "[XOH-05]")
		})
	}
}
