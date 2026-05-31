//ff:func feature=openapi-parse type=test control=sequence
//ff:what collectItemFields/extractArrayItemFields/extractBodyConstraints/extractResponseFields/collect*ForOp/fill*Lines 직접 단위 검증
package openapi

import (
	"testing"
)

func TestCollectRequestConstraintsForOp(t *testing.T) {
	result := map[string]map[string]FieldConstraint{}
	op := jsonBodyOp("CreateUser", strProps("email"))
	collectRequestConstraintsForOp(result, op, nil)
	if result["CreateUser"]["email"].Type != "string" {
		t.Errorf("result = %v", result)
	}
	// op without operationId is skipped
	noID := jsonBodyOp("", strProps("x"))
	collectRequestConstraintsForOp(result, noID, nil)
	if _, ok := result[""]; ok {
		t.Errorf("empty opID should be skipped")
	}
}
