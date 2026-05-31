//ff:func feature=openapi-parse type=test control=sequence
//ff:what collectItemFields/extractArrayItemFields/extractBodyConstraints/extractResponseFields/collect*ForOp/fill*Lines 직접 단위 검증
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractResponseFields(t *testing.T) {
	op := jsonRespOp("x", strProps("token"))
	fc := extractResponseFields(op)
	if len(fc) != 1 || fc["token"].Type != "string" {
		t.Errorf("response fields = %v", fc)
	}
	// no 2xx
	op2 := openapi3.NewOperation()
	op2.OperationID = "y"
	r := openapi3.NewResponses()
	r.Set("404", &openapi3.ResponseRef{Value: openapi3.NewResponse()})
	op2.Responses = r
	if extractResponseFields(op2) != nil {
		t.Errorf("no 2xx should yield nil")
	}
}
