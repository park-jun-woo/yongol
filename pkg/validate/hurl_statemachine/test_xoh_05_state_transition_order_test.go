//ff:func feature=validate type=test control=sequence topic=hurl-statemachine
//ff:what XOH-05 positive/negative — state 전이 순서

package hurl_statemachine

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// newWorkflowDoc returns a minimal OpenAPI doc with the zenflow-like
// workflow endpoints used across the tests in this file.
func newWorkflowDoc() *openapi3.T {
	doc := &openapi3.T{Paths: &openapi3.Paths{}}
	post := &openapi3.PathItem{Post: &openapi3.Operation{OperationID: "CreateWorkflow", Responses: openapi3.NewResponses()}}
	doc.Paths.Set("/workflows", post)
	activate := &openapi3.PathItem{Post: &openapi3.Operation{OperationID: "ActivateWorkflow", Responses: openapi3.NewResponses()}}
	doc.Paths.Set("/workflows/{id}/activate", activate)
	execute := &openapi3.PathItem{Post: &openapi3.Operation{OperationID: "ExecuteWorkflow", Responses: openapi3.NewResponses()}}
	doc.Paths.Set("/workflows/{id}/execute", execute)
	return doc
}

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

func TestXoh05_Negative_ExecuteBeforeActivate(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc:    newWorkflowDoc(),
		StateDiagrams: []*statemachine.StateDiagram{newWorkflowDiagram()},
		HurlEntries: []hurl.HurlEntry{
			{Method: "POST", Path: "/workflows", File: "t.hurl", Line: 1},
			{Method: "POST", Path: "/workflows/1/execute", File: "t.hurl", Line: 5},
		},
	}
	diags := xoh05StateTransitionOrder(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XOH-05]") || !strings.Contains(diags[0].Message, "ExecuteWorkflow") {
		t.Fatalf("unexpected msg: %q", diags[0].Message)
	}
}

func TestXoh05_Positive_ActivateThenExecute(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc:    newWorkflowDoc(),
		StateDiagrams: []*statemachine.StateDiagram{newWorkflowDiagram()},
		HurlEntries: []hurl.HurlEntry{
			{Method: "POST", Path: "/workflows", File: "t.hurl", Line: 1},
			{Method: "POST", Path: "/workflows/1/activate", File: "t.hurl", Line: 3},
			{Method: "POST", Path: "/workflows/1/execute", File: "t.hurl", Line: 5},
		},
	}
	if diags := xoh05StateTransitionOrder(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %+v", diags)
	}
}
