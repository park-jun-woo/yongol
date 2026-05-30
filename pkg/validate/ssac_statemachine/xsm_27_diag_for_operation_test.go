//ff:func feature=validate type=test control=selection topic=states
//ff:what TestXsm27DiagForOperation — xsm27DiagForOperation 단일 operation XSM-27 gate 분기 검증

package ssac_statemachine

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestXsm27DiagForOperation(t *testing.T) {
	target := &statefulTarget{
		Resource: "order",
		Model:    "Order",
		Diagram:  &statemachine.StateDiagram{ID: "order", InitialState: "draft"},
	}
	findByID := ssac.Sequence{Type: "get", Model: "Order.FindByID", Result: &ssac.Result{Var: "order"}}

	t.Run("nil op", func(t *testing.T) {
		if _, ok := xsm27DiagForOperation("POST", nil, target, nil); ok {
			t.Error("expected ok=false")
		}
	})

	t.Run("empty operation id", func(t *testing.T) {
		op := &openapi3.Operation{OperationID: ""}
		if _, ok := xsm27DiagForOperation("POST", op, target, nil); ok {
			t.Error("expected ok=false")
		}
	})

	t.Run("func not found", func(t *testing.T) {
		op := &openapi3.Operation{OperationID: "CancelOrder"}
		if _, ok := xsm27DiagForOperation("POST", op, target, map[string]ssac.ServiceFunc{}); ok {
			t.Error("expected ok=false")
		}
	})

	t.Run("state neutral", func(t *testing.T) {
		op := &openapi3.Operation{OperationID: "CancelOrder"}
		fbn := map[string]ssac.ServiceFunc{"CancelOrder": {Name: "CancelOrder", StateNeutral: true}}
		if _, ok := xsm27DiagForOperation("POST", op, target, fbn); ok {
			t.Error("expected ok=false")
		}
	})

	t.Run("has state sequence", func(t *testing.T) {
		op := &openapi3.Operation{OperationID: "CancelOrder"}
		fbn := map[string]ssac.ServiceFunc{"CancelOrder": {
			Name:      "CancelOrder",
			Sequences: []ssac.Sequence{{Type: "state"}},
		}}
		if _, ok := xsm27DiagForOperation("POST", op, target, fbn); ok {
			t.Error("expected ok=false")
		}
	})

	t.Run("resource not read via FindByID", func(t *testing.T) {
		op := &openapi3.Operation{OperationID: "CancelOrder"}
		fbn := map[string]ssac.ServiceFunc{"CancelOrder": {
			Name:      "CancelOrder",
			Sequences: []ssac.Sequence{{Type: "post"}},
		}}
		if _, ok := xsm27DiagForOperation("POST", op, target, fbn); ok {
			t.Error("expected ok=false")
		}
	})

	t.Run("all gates pass", func(t *testing.T) {
		op := &openapi3.Operation{OperationID: "CancelOrder"}
		fbn := map[string]ssac.ServiceFunc{"CancelOrder": {
			Name:      "CancelOrder",
			FileName:  "service/order/cancel.ssac",
			Sequences: []ssac.Sequence{findByID},
		}}
		d, ok := xsm27DiagForOperation("POST", op, target, fbn)
		if !ok {
			t.Fatal("expected ok=true")
		}
		if d.OperationID == "" && d.Message == "" {
			t.Errorf("expected populated diagnostic, got %+v", d)
		}
	})
}
