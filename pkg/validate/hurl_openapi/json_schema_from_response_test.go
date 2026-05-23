//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what jsonSchemaFromResponse — response에서 첫 JSON 미디어 타입 schema 추출 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestJsonSchemaFromResponse(t *testing.T) {
	targetSchema := &openapi3.Schema{Type: &openapi3.Types{"object"}}

	cases := []struct {
		name    string
		ref     *openapi3.ResponseRef
		wantNil bool
		wantPtr *openapi3.Schema
	}{
		{name: "nil_ref", ref: nil, wantNil: true},
		{name: "nil_value", ref: &openapi3.ResponseRef{Value: nil}, wantNil: true},
		{
			name: "no_content",
			ref: &openapi3.ResponseRef{
				Value: &openapi3.Response{},
			},
			wantNil: true,
		},
		{
			name: "json_content",
			ref: &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{Value: targetSchema},
						},
					},
				},
			},
			wantPtr: targetSchema,
		},
		{
			name: "xml_content_no_json",
			ref: &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"application/xml": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{Value: targetSchema},
						},
					},
				},
			},
			wantNil: true,
		},
		{
			name: "json_content_nil_schema",
			ref: &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{Schema: nil},
					},
				},
			},
			wantNil: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runSchemaPointerCase(t, jsonSchemaFromResponse(c.ref), c.wantNil, c.wantPtr)
		})
	}
}
