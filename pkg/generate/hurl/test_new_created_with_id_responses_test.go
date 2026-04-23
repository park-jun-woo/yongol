//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what newCreatedWithIDResponses — id 필드 schema 포함 201 Responses 생성 (workflow_id capture 용)

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// newCreatedWithIDResponses returns a 201 response whose schema carries
// an `id` field. Used for resource-creating operations so the smoke
// walker captures `{resource}_id` for subsequent path substitution.
func newCreatedWithIDResponses() *openapi3.Responses {
	desc := "created"
	schema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"id": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}
	r := openapi3.NewResponses()
	r.Set("201", &openapi3.ResponseRef{Value: openapi3.NewResponse().
		WithDescription(desc).
		WithContent(openapi3.NewContentWithJSONSchema(schema))})
	return r
}
