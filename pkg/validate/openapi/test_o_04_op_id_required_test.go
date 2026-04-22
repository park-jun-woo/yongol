//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-4 테스트 — all-present / one-missing / multi-missing / empty paths

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func newFullstackWithPaths(paths *openapi3.Paths) *yongol.Fullstack {
	return &yongol.Fullstack{
		OpenAPIDoc: &openapi3.T{Paths: paths},
		OpenAPILines: &oapiparser.LineIndex{
			Paths: map[string]int{},
		},
	}
}

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

func TestO04OpIdRequired_MultiMissing(t *testing.T) {
	paths := openapi3.NewPaths(
		openapi3.WithPath("/a", &openapi3.PathItem{
			Get:  &openapi3.Operation{}, // missing
			Post: &openapi3.Operation{OperationID: "CreateA"},
		}),
		openapi3.WithPath("/b", &openapi3.PathItem{
			Get:    &openapi3.Operation{OperationID: "GetB"},
			Delete: &openapi3.Operation{}, // missing
		}),
	)
	fs := newFullstackWithPaths(paths)

	diags := o04OpIdRequired(fs)
	if len(diags) != 2 {
		t.Fatalf("expected 2 diagnostics, got %d: %+v", len(diags), diags)
	}

	var sawA, sawB bool
	for _, d := range diags {
		if strings.Contains(d.Message, "GET /a") {
			sawA = true
		}
		if strings.Contains(d.Message, "DELETE /b") {
			sawB = true
		}
	}
	if !sawA || !sawB {
		t.Errorf("expected both GET /a and DELETE /b diagnostics, got %+v", diags)
	}
}

func TestO04OpIdRequired_NilPaths(t *testing.T) {
	fs := &yongol.Fullstack{OpenAPIDoc: &openapi3.T{}}
	diags := o04OpIdRequired(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics for nil Paths, got %d: %+v", len(diags), diags)
	}

	fs2 := &yongol.Fullstack{}
	if diags := o04OpIdRequired(fs2); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics for nil OpenAPIDoc, got %d", len(diags))
	}
}
