//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what operationIDSet — nil/empty/복수 operation 수집 검증

package openapi_manifest

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestOperationIDSet(t *testing.T) {
	t.Run("nil fullstack returns empty", func(t *testing.T) {
		got := operationIDSet(nil)
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("nil doc returns empty", func(t *testing.T) {
		got := operationIDSet(&yongol.Fullstack{})
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("collects operationIds", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users", &openapi3.PathItem{
					Get:  &openapi3.Operation{OperationID: "listUsers"},
					Post: &openapi3.Operation{OperationID: "createUser"},
				}),
				openapi3.WithPath("/users/{id}", &openapi3.PathItem{
					Get: &openapi3.Operation{OperationID: "getUser"},
				}),
			),
		}
		fs := &yongol.Fullstack{OpenAPIDoc: doc}
		got := operationIDSet(fs)
		if len(got) != 3 {
			t.Fatalf("expected 3, got %d: %v", len(got), got)
		}
		if !got["listUsers"] || !got["createUser"] || !got["getUser"] {
			t.Errorf("missing expected IDs: %v", got)
		}
	})

	t.Run("empty operationId skipped", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/health", &openapi3.PathItem{
					Get: &openapi3.Operation{},
				}),
			),
		}
		fs := &yongol.Fullstack{OpenAPIDoc: doc}
		got := operationIDSet(fs)
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})
}
