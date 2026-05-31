//ff:func feature=rule type=test control=iteration dimension=1
//ff:what registerOpenAPIResponseProps test — 2xx schema properties → OpenAPI.response.<op>.<field> 등록 (nil guard·맥락·skip)
package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestRegisterOpenAPIResponseProps(t *testing.T) {
	const op = "Op"
	const key = "OpenAPI.response.Op."

	schema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			// CONTEXT divergence: response date-time → time.Time (param → string).
			"createdAt": strSchema("date-time"),
			"id":        strSchema("uuid"),
			"name":      strSchema(""),
			// resolves to "" (untyped) → must be skipped, key absent.
			"untyped": {Value: &openapi3.Schema{}},
		},
	}

	g := newGround()
	registerOpenAPIResponseProps(g, op, schema)

	want := map[string]string{
		"createdAt": "time.Time",
		"id":        "openapi_types.UUID",
		"name":      "string",
	}
	for field, wantType := range want {
		if got := g.Types[key+field]; got != wantType {
			t.Errorf("%s%s = %q, want %q", key, field, got, wantType)
		}
	}
	if _, found := g.Types[key+"untyped"]; found {
		t.Errorf("%suntyped must be skipped (resolves to \"\"), but was registered", key)
	}
}
