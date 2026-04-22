//ff:func feature=rule type=test control=sequence dimension=1
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
