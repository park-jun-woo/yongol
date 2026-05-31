//ff:func feature=openapi-parse type=test control=sequence
//ff:what collectItemFields/extractArrayItemFields/extractBodyConstraints/extractResponseFields/collect*ForOp/fill*Lines 직접 단위 검증
package openapi

import (
	"testing"
)

func TestCollectResponseConstraintsForOp(t *testing.T) {
	result := map[string]map[string]FieldConstraint{}
	op := jsonRespOp("GetUser", strProps("id"))
	collectResponseConstraintsForOp(result, op, nil)
	if result["GetUser"]["id"].Type != "string" {
		t.Errorf("result = %v", result)
	}
}
