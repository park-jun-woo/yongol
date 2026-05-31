//ff:func feature=generate type=test control=sequence
//ff:what TestGenerateBatch_ZeroCov — pkg/generate 소형 순수/IO 헬퍼 분기 커버
package generate

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractRequestBodySchema_ZeroCov(t *testing.T) {
	// nil request body
	if extractRequestBodySchema(&openapi3.Operation{}) != nil {
		t.Error("nil body should return nil")
	}
	// application/json present
	jsonSchema := &openapi3.Schema{Type: &openapi3.Types{"object"}}
	op := &openapi3.Operation{RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
		Content: openapi3.Content{
			"application/json": &openapi3.MediaType{Schema: &openapi3.SchemaRef{Value: jsonSchema}},
		},
	}}}
	if extractRequestBodySchema(op) != jsonSchema {
		t.Error("should return json schema")
	}
	// fallback content type
	xmlSchema := &openapi3.Schema{Type: &openapi3.Types{"string"}}
	op2 := &openapi3.Operation{RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
		Content: openapi3.Content{
			"application/xml": &openapi3.MediaType{Schema: &openapi3.SchemaRef{Value: xmlSchema}},
		},
	}}}
	if extractRequestBodySchema(op2) != xmlSchema {
		t.Error("should fall back to xml schema")
	}
	// content present but no usable schema
	op3 := &openapi3.Operation{RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
		Content: openapi3.Content{"application/json": &openapi3.MediaType{}},
	}}}
	if extractRequestBodySchema(op3) != nil {
		t.Error("no schema → nil")
	}
}
