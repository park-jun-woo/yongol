//ff:func feature=rule type=test control=sequence
//ff:what populateResponseSchema 4xx/5xx 케이스 — primary 2xx 키는 4xx/5xx에 의해 오염되지 않는다

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestPopulateResponseSchema_4xx5xx_NoPrimary2xx verifies that when only 4xx/5xx
// responses exist, code-specific keys are populated but the back-compat
// OpenAPI.response.<opID> key (2xx-only) must NOT be written.
func TestPopulateResponseSchema_4xx5xx_NoPrimary2xx(t *testing.T) {
	op := &openapi3.Operation{OperationID: "ErrOnly", Responses: openapi3.NewResponses()}
	setJSONResponse(op, "400", []string{"message"})
	setJSONResponse(op, "500", []string{"error"})

	g := newGround()
	populateResponseSchema(g, "ErrOnly", op)

	if _, ok := g.Schemas["OpenAPI.response.400.ErrOnly"]; !ok {
		t.Errorf("400 key missing")
	}
	if _, ok := g.Schemas["OpenAPI.response.500.ErrOnly"]; !ok {
		t.Errorf("500 key missing")
	}
	if _, ok := g.Schemas["OpenAPI.response.ErrOnly"]; ok {
		t.Errorf("back-compat primary key must not be written when no 2xx present")
	}
}
