//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what requestBodyAllRequired — 전필드 required true / 일부 optional false / requestBody 부재 false 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestRequestBodyAllRequired(t *testing.T) {
	allReq := postOp("Op", map[string]*openapi3.SchemaRef{"a": stringProp(), "b": intProp()})
	allReq.Post.RequestBody.Value.Content["application/json"].Schema.Value.Required = []string{"a", "b"}
	if !requestBodyAllRequired(allReq.Post) {
		t.Errorf("all-required body should report true")
	}

	partial := postOp("Op", map[string]*openapi3.SchemaRef{"a": stringProp(), "b": intProp()})
	partial.Post.RequestBody.Value.Content["application/json"].Schema.Value.Required = []string{"a"}
	if requestBodyAllRequired(partial.Post) {
		t.Errorf("partial-required body should report false")
	}

	// No requestBody.
	noBody := &openapi3.Operation{OperationID: "Op", Responses: openapi3.NewResponses()}
	if requestBodyAllRequired(noBody) {
		t.Errorf("no requestBody should report false")
	}

	// Empty properties.
	empty := postOp("Op", map[string]*openapi3.SchemaRef{})
	if requestBodyAllRequired(empty.Post) {
		t.Errorf("empty body should report false")
	}

	if requestBodyAllRequired(nil) {
		t.Errorf("nil op should report false")
	}
}
