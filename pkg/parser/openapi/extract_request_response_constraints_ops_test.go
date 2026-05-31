//ff:func feature=openapi-parse type=test control=sequence
//ff:what collectItemFields/extractArrayItemFields/extractBodyConstraints/extractResponseFields/collect*ForOp/fill*Lines 직접 단위 검증
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractRequestResponseConstraintsOps(t *testing.T) {
	item := &openapi3.PathItem{Post: jsonBodyOp("CreateUser", strProps("email"))}
	result := map[string]map[string]FieldConstraint{}
	extractRequestConstraintsOps(result, item, nil)
	if result["CreateUser"]["email"].Type != "string" {
		t.Errorf("request ops result = %v", result)
	}

	item2 := &openapi3.PathItem{Get: jsonRespOp("GetUser", strProps("id"))}
	result2 := map[string]map[string]FieldConstraint{}
	extractResponseConstraintsOps(result2, item2, nil)
	if result2["GetUser"]["id"].Type != "string" {
		t.Errorf("response ops result = %v", result2)
	}
}
