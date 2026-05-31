//ff:func feature=openapi-parse type=test control=iteration dimension=1
//ff:what TestOpenAPIHelpers — unit tests for the pure openapi parser helper functions
package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func opWith2xx(codes ...int) *openapi3.Operation {
	op := &openapi3.Operation{Responses: openapi3.NewResponses()}
	for _, c := range codes {
		op.Responses.Set(itoaCode(c), &openapi3.ResponseRef{Value: &openapi3.Response{}})
	}
	return op
}
