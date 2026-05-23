//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what buildOperationMethodMap — nil/empty/method 포함 수집 검증

package openapi_ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildOperationMethodMap(t *testing.T) {
	t.Run("nil doc returns empty", func(t *testing.T) {
		got := buildOperationMethodMap(nil)
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("collects with method", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users", &openapi3.PathItem{
					Get:  &openapi3.Operation{OperationID: "listUsers"},
					Post: &openapi3.Operation{OperationID: "createUser"},
				}),
			),
		}
		got := buildOperationMethodMap(doc)
		if len(got) != 2 {
			t.Fatalf("expected 2, got %d", len(got))
		}
		if got["listUsers"].Method != "GET" {
			t.Errorf("listUsers method = %q, want GET", got["listUsers"].Method)
		}
		if got["createUser"].Method != "POST" {
			t.Errorf("createUser method = %q, want POST", got["createUser"].Method)
		}
	})
}
