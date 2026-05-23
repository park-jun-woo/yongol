//ff:func feature=validate type=test control=sequence topic=hurl-statemachine
//ff:what inspectDiagramEntry — entry가 diagram 전이에 해당하는지 보고 reachable 갱신 검증

package hurl_statemachine

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestInspectDiagramEntry(t *testing.T) {
	d := &statemachine.StateDiagram{
		ID: "order",
		Transitions: []statemachine.Transition{
			{From: "draft", To: "submitted", Event: "submitOrder"},
			{From: "submitted", To: "approved", Event: "approveOrder"},
		},
	}
	opID := map[string]string{
		"POST /orders": "submitOrder",
	}

	t.Run("unrelated_entry_no_diag", func(t *testing.T) {
		reachable := map[string]bool{"draft": true}
		e := hurl.HurlEntry{Method: "GET", Path: "/users"}
		diags := inspectDiagramEntry(e, opID, d, reachable)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("reachable_from_state_advances", func(t *testing.T) {
		reachable := map[string]bool{"draft": true}
		e := hurl.HurlEntry{Method: "POST", Path: "/orders"}
		diags := inspectDiagramEntry(e, opID, d, reachable)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
		if !reachable["submitted"] {
			t.Errorf("expected 'submitted' in reachable")
		}
	})

	t.Run("op_not_in_diagram_no_diag", func(t *testing.T) {
		// opID maps the entry, but the operationId is not in this diagram's transitions
		opIDExtra := map[string]string{"GET /health": "healthCheck"}
		reachable := map[string]bool{"draft": true}
		e := hurl.HurlEntry{Method: "GET", Path: "/health"}
		diags := inspectDiagramEntry(e, opIDExtra, d, reachable)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("unreachable_from_state_produces_warning", func(t *testing.T) {
		reachable := map[string]bool{"approved": true} // draft not reachable
		e := hurl.HurlEntry{Method: "POST", Path: "/orders", File: "t.hurl", Line: 1}
		diags := inspectDiagramEntry(e, opID, d, reachable)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[XOH-05]") {
			t.Errorf("expected [XOH-05], got %q", diags[0].Message)
		}
	})
}
