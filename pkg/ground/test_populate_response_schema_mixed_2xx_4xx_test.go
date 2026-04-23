//ff:func feature=rule type=test control=sequence
//ff:what populateResponseSchema — 2xx+4xx 혼합 시 primary 2xx 키는 첫 2xx 로 고정

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestPopulateResponseSchema_Mixed2xx4xx verifies the primary 2xx key is set
// to the FIRST 2xx content (not overwritten by later 4xx).
func TestPopulateResponseSchema_Mixed2xx4xx(t *testing.T) {
	op := &openapi3.Operation{OperationID: "Mixed", Responses: openapi3.NewResponses()}
	setJSONResponse(op, "200", []string{"a", "b"})
	setJSONResponse(op, "404", []string{"message"})

	g := newGround()
	populateResponseSchema(g, "Mixed", op)

	primary, ok := g.Schemas["OpenAPI.response.Mixed"]
	if !ok {
		t.Fatalf("primary 2xx key missing")
	}
	if len(primary) != 2 {
		t.Errorf("primary 2xx fields len = %d, want 2", len(primary))
	}

	err404, ok := g.Schemas["OpenAPI.response.404.Mixed"]
	if !ok {
		t.Fatalf("404 code key missing")
	}
	if len(err404) != 1 {
		t.Errorf("404 fields len = %d, want 1", len(err404))
	}
}
