//ff:func feature=openapi-parse type=test control=iteration dimension=1
//ff:what collectItemFields/extractArrayItemFields/extractBodyConstraints/extractResponseFields/collect*ForOp/fill*Lines 직접 단위 검증
package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func arraySchemaRef(itemProps ...string) *openapi3.SchemaRef {
	props := openapi3.Schemas{}
	for _, p := range itemProps {
		props[p] = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	}
	return &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{Value: &openapi3.Schema{Properties: props}},
	}}
}
