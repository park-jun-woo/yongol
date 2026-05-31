//ff:func feature=openapi-parse type=test control=sequence
//ff:what collectItemFields/extractArrayItemFields/extractBodyConstraints/extractResponseFields/collect*ForOp/fill*Lines 직접 단위 검증
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractBodyConstraints(t *testing.T) {
	op := jsonBodyOp("x", strProps("email", "name"))
	fc := extractBodyConstraints(op.RequestBody, "x")
	if len(fc) != 2 || fc["email"].Type != "string" {
		t.Errorf("constraints = %v", fc)
	}
	// nil content body
	empty := &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody()}
	if extractBodyConstraints(empty, "x") != nil {
		t.Errorf("empty body should yield nil")
	}
}
