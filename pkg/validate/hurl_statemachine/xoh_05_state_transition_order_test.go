//ff:func feature=validate type=test control=sequence topic=hurl-statemachine
//ff:what xoh05StateTransitionOrder — 파일 내 operation 호출 순서가 state machine 전이 규칙 준수 검증

package hurl_statemachine

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh05StateTransitionOrder(t *testing.T) {
	t.Run("nil_fs", func(t *testing.T) {
		diags := xoh05StateTransitionOrder(nil)
		if len(diags) != 0 {
			t.Errorf("expected 0, got %d", len(diags))
		}
	})

	t.Run("no_diagrams", func(t *testing.T) {
		fs := &yongol.Fullstack{
			HurlEntries: []hurl.HurlEntry{{Method: "GET", Path: "/users"}},
		}
		diags := xoh05StateTransitionOrder(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0, got %d", len(diags))
		}
	})

	t.Run("no_entries", func(t *testing.T) {
		fs := &yongol.Fullstack{
			StateDiagrams: []*statemachine.StateDiagram{{ID: "order"}},
		}
		diags := xoh05StateTransitionOrder(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0, got %d", len(diags))
		}
	})

	t.Run("no_operation_ids_in_doc", func(t *testing.T) {
		doc := &openapi3.T{Paths: &openapi3.Paths{}}
		doc.Paths.Set("/health", &openapi3.PathItem{
			Get: &openapi3.Operation{Responses: &openapi3.Responses{}}, // no OperationID
		})
		fs := &yongol.Fullstack{
			OpenAPIDoc: doc,
			StateDiagrams: []*statemachine.StateDiagram{{
				ID: "order", InitialState: "draft",
				Transitions: []statemachine.Transition{{From: "draft", To: "submitted", Event: "submit"}},
			}},
			HurlEntries: []hurl.HurlEntry{{Method: "GET", Path: "/health", File: "t.hurl"}},
		}
		diags := xoh05StateTransitionOrder(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0, got %d", len(diags))
		}
	})

	t.Run("correct_order_no_diag", func(t *testing.T) {
		doc := &openapi3.T{Paths: &openapi3.Paths{}}
		doc.Paths.Set("/orders", &openapi3.PathItem{
			Post: &openapi3.Operation{OperationID: "submitOrder", Responses: &openapi3.Responses{}},
		})

		fs := &yongol.Fullstack{
			OpenAPIDoc: doc,
			StateDiagrams: []*statemachine.StateDiagram{{
				ID: "order", InitialState: "draft",
				Transitions: []statemachine.Transition{
					{From: "draft", To: "submitted", Event: "submitOrder"},
				},
			}},
			HurlEntries: []hurl.HurlEntry{
				{Method: "POST", Path: "/orders", File: "t.hurl", Line: 1},
			},
		}
		diags := xoh05StateTransitionOrder(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("wrong_order_produces_warning", func(t *testing.T) {
		doc := &openapi3.T{Paths: &openapi3.Paths{}}
		doc.Paths.Set("/orders/{id}/approve", &openapi3.PathItem{
			Post: &openapi3.Operation{OperationID: "approveOrder", Responses: &openapi3.Responses{}},
		})

		fs := &yongol.Fullstack{
			OpenAPIDoc: doc,
			StateDiagrams: []*statemachine.StateDiagram{{
				ID: "order", InitialState: "draft",
				Transitions: []statemachine.Transition{
					{From: "draft", To: "submitted", Event: "submitOrder"},
					{From: "submitted", To: "approved", Event: "approveOrder"},
				},
			}},
			HurlEntries: []hurl.HurlEntry{
				{Method: "POST", Path: "/orders/1/approve", File: "t.hurl", Line: 1},
			},
		}
		diags := xoh05StateTransitionOrder(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[XOH-05]") {
			t.Errorf("expected [XOH-05], got %q", diags[0].Message)
		}
	})
}
