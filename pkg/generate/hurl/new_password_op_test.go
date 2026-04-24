//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what newPasswordOp — auth shape 테스트용 *openapi3.Operation 픽스처

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// newPasswordOp returns a minimal *openapi3.Operation for shape-
// detection tests. public=true sets Security to an empty non-nil slice
// (the OpenAPI "explicit public override"). hasPassword=true wires a
// JSON body with a "password" property.
func newPasswordOp(opID string, public, hasPassword bool) *openapi3.Operation {
	op := &openapi3.Operation{OperationID: opID}
	if public {
		emptySec := openapi3.SecurityRequirements{}
		op.Security = &emptySec
	}
	if hasPassword {
		body := &openapi3.Schema{
			Type: &openapi3.Types{"object"},
			Properties: openapi3.Schemas{
				"email":    {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
				"password": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			},
		}
		op.RequestBody = &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().
			WithContent(openapi3.NewContentWithJSONSchema(body))}
	}
	op.Responses = newOKResponses()
	return op
}
