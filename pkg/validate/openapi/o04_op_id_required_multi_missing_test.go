//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-structural
//ff:what O-4 — 여러 operation 누락 시 각 진단 발화

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

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
