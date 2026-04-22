//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateOpenAPIParams — operationId 부재 operation은 skip (Ground 계층은 diagnostic 미발행)

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestPopulateOpenAPIParams_MissingOpID verifies that operations with empty
// OperationID are skipped silently (O-4 upstream gates this).
func TestPopulateOpenAPIParams_MissingOpID(t *testing.T) {
	param := &openapi3.ParameterRef{Value: &openapi3.Parameter{Name: "q", In: "query"}}
	op := &openapi3.Operation{
		Parameters: openapi3.Parameters{param},
	}
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/noop", &openapi3.PathItem{Get: op}),
	)}
	fs := newMinimalFullstack(withOpenAPIDoc(doc))
	g := newGround()

	populateOpenAPIParams(g, fs)

	for k := range g.Lookup {
		// No OpenAPI.param.* key should exist, because we have no opID to key on.
		if len(k) >= len("OpenAPI.param.") && k[:len("OpenAPI.param.")] == "OpenAPI.param." {
			t.Errorf("unexpected key %q when operationId is missing", k)
		}
	}
}

// TestPopulateOpenAPIParams_NilDoc: nil doc must be tolerated.
func TestPopulateOpenAPIParams_NilDoc(t *testing.T) {
	g := newGround()
	populateOpenAPIParams(g, newMinimalFullstack())
	if len(g.Lookup) != 0 {
		t.Errorf("expected 0 keys, got %d", len(g.Lookup))
	}
}
