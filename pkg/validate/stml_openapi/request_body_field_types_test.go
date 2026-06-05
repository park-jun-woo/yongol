//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what requestBodyFieldTypes — requestBody 스키마에서 property→type 추출 분기 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestRequestBodyFieldTypes(t *testing.T) {
	// nil op → empty.
	if got := requestBodyFieldTypes(nil); len(got) != 0 {
		t.Fatalf("nil op: expected empty, got %+v", got)
	}
	// op with nil RequestBody → empty.
	if got := requestBodyFieldTypes(&openapi3.Operation{}); len(got) != 0 {
		t.Fatalf("nil body: expected empty, got %+v", got)
	}
	// nil RequestBody.Value → empty.
	op := &openapi3.Operation{RequestBody: &openapi3.RequestBodyRef{Value: nil}}
	if got := requestBodyFieldTypes(op); len(got) != 0 {
		t.Fatalf("nil body value: expected empty, got %+v", got)
	}

	// Content with only a nil-schema media type → continue, then return empty.
	opNilSchema := &openapi3.Operation{
		RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{Schema: nil},
			},
		}},
	}
	if got := requestBodyFieldTypes(opNilSchema); len(got) != 0 {
		t.Fatalf("nil schema: expected empty, got %+v", got)
	}

	// Valid request body with typed properties.
	opValid := &openapi3.Operation{
		RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
						Properties: openapi3.Schemas{
							"title": typedSchema("string"),
							"count": typedSchema("integer"),
						},
					}},
				},
			},
		}},
	}
	out := requestBodyFieldTypes(opValid)
	want := map[string]string{"title": "string", "count": "integer"}
	if len(out) != len(want) {
		t.Fatalf("got %+v, want %+v", out, want)
	}
	for k, v := range want {
		if out[k] != v {
			t.Errorf("out[%q] = %q, want %q", k, out[k], v)
		}
	}
}
