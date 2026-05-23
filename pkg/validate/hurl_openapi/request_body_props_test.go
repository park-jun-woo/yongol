//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what requestBodyProps — operation requestBody의 top-level property 추출 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestRequestBodyProps(t *testing.T) {
	t.Run("nil_op", func(t *testing.T) {
		props, ok := requestBodyProps(nil)
		if ok || props != nil {
			t.Errorf("expected (nil, false), got (%v, %v)", props, ok)
		}
	})

	t.Run("nil_request_body", func(t *testing.T) {
		props, ok := requestBodyProps(&openapi3.Operation{})
		if ok || props != nil {
			t.Errorf("expected (nil, false), got (%v, %v)", props, ok)
		}
	})

	t.Run("json_body_with_properties", func(t *testing.T) {
		op := &openapi3.Operation{
			RequestBody: &openapi3.RequestBodyRef{
				Value: openapi3.NewRequestBody().WithJSONSchema(&openapi3.Schema{
					Properties: openapi3.Schemas{
						"name":  &openapi3.SchemaRef{Value: &openapi3.Schema{}},
						"email": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
					},
				}),
			},
		}
		props, ok := requestBodyProps(op)
		if !ok {
			t.Fatal("expected ok=true")
		}
		if len(props) != 2 {
			t.Fatalf("expected 2 props, got %d", len(props))
		}
		if _, has := props["name"]; !has {
			t.Error("missing 'name'")
		}
		if _, has := props["email"]; !has {
			t.Error("missing 'email'")
		}
	})

	t.Run("nil_schema_in_content", func(t *testing.T) {
		op := &openapi3.Operation{
			RequestBody: &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{Schema: nil},
					},
				},
			},
		}
		props, ok := requestBodyProps(op)
		if ok || props != nil {
			t.Errorf("expected (nil, false), got (%v, %v)", props, ok)
		}
	})
}
