//ff:func feature=rule type=test control=sequence
//ff:what populateOpenAPIResponseTypesSingle — operationId 누락/비2xx/비JSON content 스킵 분기 검증

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestPopulateOpenAPIResponseTypesSingle_SkipBranches exercises the reachable
// skip paths: an empty operationId, an operation with only a non-2xx response,
// and a 2xx response that carries no application/json content. None should
// register response field types.
func TestPopulateOpenAPIResponseTypesSingle_SkipBranches(t *testing.T) {
	// op with empty operationId — skipped.
	emptyID := &openapi3.Operation{OperationID: ""}
	emptyID.Responses = openapi3.NewResponses()

	// op whose only response is 4xx — the "2" prefix check skips it.
	only4xx := &openapi3.Operation{OperationID: "Only4xx"}
	resp4 := openapi3.NewResponse().WithContent(openapi3.NewContentWithJSONSchema(&openapi3.Schema{Type: &openapi3.Types{"object"}}))
	only4xx.Responses = openapi3.NewResponses()
	only4xx.Responses.Set("404", &openapi3.ResponseRef{Value: resp4})

	// op with a 2xx response but no application/json content — ct nil skip.
	noJSON := &openapi3.Operation{OperationID: "NoJSON"}
	respNoJSON := openapi3.NewResponse()
	respNoJSON.Content = openapi3.Content{} // present but no application/json
	noJSON.Responses = openapi3.NewResponses()
	noJSON.Responses.Set("200", &openapi3.ResponseRef{Value: respNoJSON})

	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/a", &openapi3.PathItem{Get: emptyID}),
		openapi3.WithPath("/b", &openapi3.PathItem{Get: only4xx}),
		openapi3.WithPath("/c", &openapi3.PathItem{Get: noJSON}),
	)}
	g := newGround()

	populateOpenAPIResponseTypesSingle(g, doc)

	if len(g.Types) != 0 {
		t.Fatalf("expected no response types registered, got %v", g.Types)
	}
}
