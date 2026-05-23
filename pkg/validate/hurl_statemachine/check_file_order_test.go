//ff:func feature=validate type=test control=sequence topic=hurl-statemachine
//ff:what checkFileOrder — 한 파일의 entries를 각 diagram마다 전이 순서 검증

package hurl_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestCheckFileOrder(t *testing.T) {
	d1 := &statemachine.StateDiagram{
		ID: "order", InitialState: "draft",
		Transitions: []statemachine.Transition{
			{From: "draft", To: "submitted", Event: "submitOrder"},
		},
	}
	d2 := &statemachine.StateDiagram{
		ID: "payment", InitialState: "pending",
		Transitions: []statemachine.Transition{
			{From: "pending", To: "paid", Event: "pay"},
		},
	}
	opID := map[string]string{
		"POST /orders": "submitOrder",
		"POST /pay":    "pay",
	}

	t.Run("no_diagrams_no_diag", func(t *testing.T) {
		diags := checkFileOrder(nil, opID, nil)
		if len(diags) != 0 {
			t.Errorf("expected 0, got %d", len(diags))
		}
	})

	t.Run("checks_multiple_diagrams", func(t *testing.T) {
		entries := []hurl.HurlEntry{
			{Method: "POST", Path: "/orders", File: "t.hurl", Line: 1},
			{Method: "POST", Path: "/pay", File: "t.hurl", Line: 5},
		}
		diags := checkFileOrder(entries, opID, []*statemachine.StateDiagram{d1, d2})
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})
}
