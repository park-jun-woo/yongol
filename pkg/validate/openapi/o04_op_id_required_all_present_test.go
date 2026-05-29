//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-4 — 모든 operationId 가 선언된 경우 진단 없음

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestO04OpIdRequired_AllPresent(t *testing.T) {
	paths := openapi3.NewPaths(
		openapi3.WithPath("/workflows", &openapi3.PathItem{
			Get:  &openapi3.Operation{OperationID: "ListWorkflows"},
			Post: &openapi3.Operation{OperationID: "CreateWorkflow"},
		}),
		openapi3.WithPath("/workflows/{id}", &openapi3.PathItem{
			Get: &openapi3.Operation{OperationID: "GetWorkflow"},
		}),
	)
	fs := newFullstackWithPaths(paths)

	diags := o04OpIdRequired(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
