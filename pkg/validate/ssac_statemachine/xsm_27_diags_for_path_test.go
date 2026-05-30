//ff:func feature=validate type=test control=selection topic=states
//ff:what TestXsm27DiagsForPath — xsm27DiagsForPath 단일 path XSM-27 진단 수집 분기 검증

package ssac_statemachine

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestXsm27DiagsForPath(t *testing.T) {
	diagram := &statemachine.StateDiagram{ID: "workflow", InitialState: "draft"}
	diagrams := []*statemachine.StateDiagram{diagram}
	ground := func() *rule.Ground {
		return &rule.Ground{Types: map[string]string{"DDL.default.value.workflows.status": "draft"}}
	}
	findByID := ssac.Sequence{Type: "get", Model: "Workflow.FindByID", Result: &ssac.Result{Var: "workflow"}}

	t.Run("nil item", func(t *testing.T) {
		if d := xsm27DiagsForPath("/workflows/{id}", nil, diagrams, ground(), nil); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("no id param", func(t *testing.T) {
		item := &openapi3.PathItem{Post: &openapi3.Operation{OperationID: "X"}}
		if d := xsm27DiagsForPath("/workflows", item, diagrams, ground(), nil); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("not stateful", func(t *testing.T) {
		item := &openapi3.PathItem{Post: &openapi3.Operation{OperationID: "X"}}
		if d := xsm27DiagsForPath("/orders/{id}", item, diagrams, ground(), nil); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("stateful path yields diagnostic", func(t *testing.T) {
		item := &openapi3.PathItem{Post: &openapi3.Operation{OperationID: "CancelWorkflow"}}
		fbn := map[string]ssac.ServiceFunc{"CancelWorkflow": {
			Name:      "CancelWorkflow",
			Sequences: []ssac.Sequence{findByID},
		}}
		diags := xsm27DiagsForPath("/workflows/{id}", item, diagrams, ground(), fbn)
		if len(diags) != 1 {
			t.Errorf("expected 1 diagnostic, got %d", len(diags))
		}
	})
}
