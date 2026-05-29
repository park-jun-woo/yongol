//ff:func feature=rule type=test control=sequence
//ff:what populateResponseSchema — Responses 가 nil 이면 조기 반환

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestPopulateResponseSchema_NilResponses covers the nil Responses early return.
func TestPopulateResponseSchema_NilResponses(t *testing.T) {
	g := newGround()
	op := &openapi3.Operation{OperationID: "NoResp"}
	populateResponseSchema(g, "NoResp", op) // must not panic

	if len(g.Schemas) != 0 {
		t.Errorf("no schemas expected when Responses is nil; got %d", len(g.Schemas))
	}
}
