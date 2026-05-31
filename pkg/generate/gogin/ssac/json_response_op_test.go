//ff:func feature=gen-gogin type=test control=sequence
//ff:what collectFrom200Response 단위 테스트 (직접 body $ref + property $ref 수집)
package ssac

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func jsonResponseOp(method string, status int, schema *openapi3.SchemaRef) *openapi3.Operation {
	return &openapi3.Operation{
		Responses: openapi3.NewResponses(
			openapi3.WithStatus(status, &openapi3.ResponseRef{
				Value: openapi3.NewResponse().
					WithDescription("ok").
					WithJSONSchemaRef(schema),
			}),
		),
	}
}
