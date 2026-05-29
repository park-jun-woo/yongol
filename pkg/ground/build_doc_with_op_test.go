//ff:func feature=rule type=test-helper control=selection
//ff:what buildDocWithOp — 단일 path+op+선택적 2xx JSON response 를 담은 최소 *openapi3.T 생성

package ground

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// buildDocWithOp returns a minimal *openapi3.T containing a single path with a
// single operation, optionally with a 2xx JSON response whose schema lists the
// given top-level field names as strings.
func buildDocWithOp(path, method, opID string, respFields []string) *openapi3.T {
	op := &openapi3.Operation{OperationID: opID}
	if respFields != nil {
		props := openapi3.Schemas{}
		for _, f := range respFields {
			props[f] = &openapi3.SchemaRef{Value: &openapi3.Schema{
				Type: &openapi3.Types{"string"},
			}}
		}
		resp := openapi3.NewResponse().
			WithContent(openapi3.NewContentWithJSONSchema(&openapi3.Schema{
				Type:       &openapi3.Types{"object"},
				Properties: props,
			}))
		responses := openapi3.NewResponses()
		responses.Set("200", &openapi3.ResponseRef{Value: resp})
		op.Responses = responses
	}
	item := &openapi3.PathItem{}
	switch method {
	case "GET":
		item.Get = op
	case "POST":
		item.Post = op
	case "PUT":
		item.Put = op
	case "DELETE":
		item.Delete = op
	}
	return &openapi3.T{Paths: openapi3.NewPaths(openapi3.WithPath(path, item))}
}
