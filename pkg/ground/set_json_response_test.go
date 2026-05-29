//ff:func feature=rule type=test-helper control=iteration dimension=1
//ff:what setJSONResponse — 특정 status code 에 JSON schema response 를 덮어써 등록

package ground

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// setJSONResponse overwrites the operation's response map with the given code
// and a JSON schema whose properties are the given string fields.
func setJSONResponse(op *openapi3.Operation, code string, fields []string) {
	props := openapi3.Schemas{}
	for _, f := range fields {
		props[f] = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	}
	resp := openapi3.NewResponse().WithContent(openapi3.NewContentWithJSONSchema(&openapi3.Schema{
		Type:       &openapi3.Types{"object"},
		Properties: props,
	}))
	if op.Responses == nil {
		op.Responses = openapi3.NewResponses()
	}
	op.Responses.Set(code, &openapi3.ResponseRef{Value: resp})
}
