//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestCollectResponseProps — 응답 content 스키마에서 property 수집 분기 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectResponseProps(t *testing.T) {
	// nil Content → no-op.
	out := map[string]responseFieldInfo{}
	collectResponseProps(out, &openapi3.Response{})
	if len(out) != 0 {
		t.Fatalf("nil content: expected empty, got %+v", out)
	}

	// First media type has nil schema (continue), second has a valid schema.
	out = map[string]responseFieldInfo{}
	content := openapi3.Content{
		"text/plain": nil, // nil media type → skipped
		"application/json": &openapi3.MediaType{
			Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"id": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
				},
			}},
		},
	}
	collectResponseProps(out, &openapi3.Response{Content: content})
	if out["id"].typ != "integer" {
		t.Fatalf("expected id typ integer, got %+v", out)
	}

	// Content with only a nil-schema media type → continue branch, no props.
	out = map[string]responseFieldInfo{}
	collectResponseProps(out, &openapi3.Response{Content: openapi3.Content{
		"application/json": &openapi3.MediaType{Schema: nil},
	}})
	if len(out) != 0 {
		t.Fatalf("nil schema: expected empty, got %+v", out)
	}
}
