//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what o04OpIdRequired — nil doc + operationId 유무 검증

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO04OpIdRequired(t *testing.T) {
	t.Run("nil OpenAPIDoc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := o04OpIdRequired(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("nil path item is skipped", func(t *testing.T) {
		paths := openapi3.NewPaths()
		paths.Set("/health", nil)
		doc := &openapi3.T{Paths: paths}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Paths: map[string]int{}},
		}
		diags := o04OpIdRequired(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("operation with operationId returns nil", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users", &openapi3.PathItem{
					Get: &openapi3.Operation{OperationID: "listUsers"},
				}),
			),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Paths: map[string]int{}, Operations: map[string]int{}},
		}
		diags := o04OpIdRequired(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("operation without operationId raises diagnostic", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users", &openapi3.PathItem{
					Get: &openapi3.Operation{},
				}),
			),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Paths: map[string]int{"/users": 10}},
		}
		diags := o04OpIdRequired(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "O-4") {
			t.Errorf("Message missing O-4: %s", diags[0].Message)
		}
	})
}
