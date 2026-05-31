//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what addItemOperations — nil op skip/operationId 수집 검증

package openapi_ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestAddItemOperations(t *testing.T) {
	opMap := map[string]*openapi3.Operation{}
	item := &openapi3.PathItem{
		Get:    &openapi3.Operation{OperationID: "getUser"},
		Post:   &openapi3.Operation{OperationID: "createUser"},
		Put:    nil,
		Delete: &openapi3.Operation{}, // no operationId
		Patch:  &openapi3.Operation{OperationID: "patchUser"},
	}

	addItemOperations(opMap, item)

	if len(opMap) != 3 {
		t.Fatalf("expected 3 ops, got %d: %v", len(opMap), opMap)
	}
	if _, ok := opMap["getUser"]; !ok {
		t.Error("missing getUser")
	}
	if _, ok := opMap["createUser"]; !ok {
		t.Error("missing createUser")
	}
	if _, ok := opMap["patchUser"]; !ok {
		t.Error("missing patchUser")
	}
}
