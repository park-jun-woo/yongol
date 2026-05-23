//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what buildOperationMap — nil doc/empty paths/복수 ops 수집 검증

package openapi_ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildOperationMap(t *testing.T) {
	t.Run("nil doc returns empty", func(t *testing.T) {
		got := buildOperationMap(nil)
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("nil paths returns empty", func(t *testing.T) {
		got := buildOperationMap(&openapi3.T{})
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("collects operations", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users", &openapi3.PathItem{
					Get:  &openapi3.Operation{OperationID: "listUsers"},
					Post: &openapi3.Operation{OperationID: "createUser"},
				}),
			),
		}
		got := buildOperationMap(doc)
		if len(got) != 2 {
			t.Fatalf("expected 2, got %d", len(got))
		}
	})
}
