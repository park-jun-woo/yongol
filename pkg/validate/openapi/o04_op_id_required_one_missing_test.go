//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-4 — operation 하나에 operationId 누락 시 ERROR 1개

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestO04OpIdRequired_OneMissing(t *testing.T) {
	paths := openapi3.NewPaths(
		openapi3.WithPath("/x", &openapi3.PathItem{
			Get:  &openapi3.Operation{}, // missing operationId
			Post: &openapi3.Operation{OperationID: "CreateX"},
		}),
	)
	fs := newFullstackWithPaths(paths)
	fs.OpenAPILines.Paths["/x"] = 42

	diags := o04OpIdRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	got := diags[0]
	if !strings.Contains(got.Message, "[O-4]") {
		t.Errorf("message missing rule id: %q", got.Message)
	}
	if !strings.Contains(got.Message, "GET /x") {
		t.Errorf("message missing method+path: %q", got.Message)
	}
	if got.File != "api/openapi.yaml" {
		t.Errorf("file = %q, want api/openapi.yaml", got.File)
	}
	if got.Line != 42 {
		t.Errorf("line = %d, want 42 (PathLine fallback)", got.Line)
	}
}
