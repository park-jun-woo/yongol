//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestRequestBodyFields — requestBody 스키마 top-level property 추출 분기 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestRequestBodyFields(t *testing.T) {
	// nil op → empty.
	if got := requestBodyFields(nil); len(got) != 0 {
		t.Fatalf("nil op: expected empty, got %+v", got)
	}
	// op with nil RequestBody → empty.
	if got := requestBodyFields(&openapi3.Operation{}); len(got) != 0 {
		t.Fatalf("nil body: expected empty, got %+v", got)
	}

	// Valid request body with properties.
	op := &openapi3.Operation{
		RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
						Properties: openapi3.Schemas{
							"title": {Value: &openapi3.Schema{}},
						},
					}},
				},
			},
		}},
	}
	out := requestBodyFields(op)
	if _, ok := out["title"]; !ok {
		t.Fatalf("expected title prop, got %+v", out)
	}

	// Content with only a nil-schema media type → continue, then return empty.
	opNil := &openapi3.Operation{
		RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{Schema: nil},
			},
		}},
	}
	if got := requestBodyFields(opNil); len(got) != 0 {
		t.Fatalf("nil schema: expected empty, got %+v", got)
	}
}
