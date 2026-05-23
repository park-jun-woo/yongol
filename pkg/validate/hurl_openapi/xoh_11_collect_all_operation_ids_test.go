//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what collectAllOperationIDs — OpenAPI Doc의 전체 operationId 집합 수집 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectAllOperationIDs(t *testing.T) {
	t.Run("nil_doc", func(t *testing.T) {
		ids := collectAllOperationIDs(nil)
		if len(ids) != 0 {
			t.Errorf("expected empty, got %v", ids)
		}
	})

	t.Run("nil_paths", func(t *testing.T) {
		ids := collectAllOperationIDs(&openapi3.T{})
		if len(ids) != 0 {
			t.Errorf("expected empty, got %v", ids)
		}
	})

	t.Run("collects_ids", func(t *testing.T) {
		doc := &openapi3.T{Paths: &openapi3.Paths{}}
		doc.Paths.Set("/users", &openapi3.PathItem{
			Get:  &openapi3.Operation{OperationID: "getUsers", Responses: &openapi3.Responses{}},
			Post: &openapi3.Operation{OperationID: "createUser", Responses: &openapi3.Responses{}},
		})
		doc.Paths.Set("/orders", &openapi3.PathItem{
			Get: &openapi3.Operation{OperationID: "getOrders", Responses: &openapi3.Responses{}},
		})
		ids := collectAllOperationIDs(doc)
		if len(ids) != 3 {
			t.Fatalf("expected 3 ids, got %d: %v", len(ids), ids)
		}
		for _, id := range []string{"getUsers", "createUser", "getOrders"} {
			if !ids[id] {
				t.Errorf("missing id %q", id)
			}
		}
	})
}
