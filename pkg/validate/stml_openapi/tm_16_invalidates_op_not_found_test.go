//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM16Invalidates — 미존재·비GET·정상 data-invalidates 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM16Invalidates(t *testing.T) {
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/workflows", &openapi3.PathItem{
		Get:  &openapi3.Operation{OperationID: "ListWorkflows"},
		Post: &openapi3.Operation{OperationID: "CreateWorkflow"},
	})
	opMap := buildOperationMethodMap(doc)

	// Normal: GET operationId → no diagnostics.
	ok := tm16InvalidatesOpNotFound(stml.ActionBlock{OperationID: "CreateWorkflow", Invalidates: []string{"ListWorkflows"}}, "p.html", opMap)
	if len(ok) != 0 {
		t.Errorf("GET invalidates: expected 0 diagnostics, got %d: %+v", len(ok), ok)
	}

	// Missing operationId → 1 diagnostic.
	missing := tm16InvalidatesOpNotFound(stml.ActionBlock{OperationID: "CreateWorkflow", Invalidates: []string{"NopeOp"}}, "p.html", opMap)
	if len(missing) != 1 {
		t.Fatalf("missing op: expected 1 diagnostic, got %d: %+v", len(missing), missing)
	}

	// Non-GET (POST) operationId → 1 diagnostic.
	nonGet := tm16InvalidatesOpNotFound(stml.ActionBlock{OperationID: "CreateWorkflow", Invalidates: []string{"CreateWorkflow"}}, "p.html", opMap)
	if len(nonGet) != 1 {
		t.Fatalf("non-GET op: expected 1 diagnostic, got %d: %+v", len(nonGet), nonGet)
	}

	// Empty Invalidates → no diagnostics.
	empty := tm16InvalidatesOpNotFound(stml.ActionBlock{OperationID: "CreateWorkflow"}, "p.html", opMap)
	if len(empty) != 0 {
		t.Errorf("empty invalidates: expected 0 diagnostics, got %d: %+v", len(empty), empty)
	}
}
