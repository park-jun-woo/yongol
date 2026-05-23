//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what addOpsFromPathItem — PathItem의 operationId를 ids에 추가 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestAddOpsFromPathItem(t *testing.T) {
	t.Run("adds_operation_ids", func(t *testing.T) {
		pi := &openapi3.PathItem{
			Get:  &openapi3.Operation{OperationID: "getUsers"},
			Post: &openapi3.Operation{OperationID: "createUser"},
		}
		ids := map[string]bool{}
		addOpsFromPathItem(ids, pi)
		if len(ids) != 2 {
			t.Fatalf("expected 2 ids, got %d", len(ids))
		}
		if !ids["getUsers"] || !ids["createUser"] {
			t.Errorf("missing ids: %v", ids)
		}
	})

	t.Run("empty_operation_id_skipped", func(t *testing.T) {
		pi := &openapi3.PathItem{
			Get: &openapi3.Operation{OperationID: ""},
		}
		ids := map[string]bool{}
		addOpsFromPathItem(ids, pi)
		if len(ids) != 0 {
			t.Errorf("expected empty, got %v", ids)
		}
	})

	t.Run("preserves_existing_ids", func(t *testing.T) {
		pi := &openapi3.PathItem{
			Get: &openapi3.Operation{OperationID: "getUsers"},
		}
		ids := map[string]bool{"existing": true}
		addOpsFromPathItem(ids, pi)
		if len(ids) != 2 || !ids["existing"] || !ids["getUsers"] {
			t.Errorf("ids = %v", ids)
		}
	})
}
