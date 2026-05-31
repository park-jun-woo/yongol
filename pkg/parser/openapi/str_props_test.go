//ff:func feature=openapi-parse type=test control=iteration dimension=1
//ff:what collectItemFields/extractArrayItemFields/extractBodyConstraints/extractResponseFields/collect*ForOp/fill*Lines 직접 단위 검증
package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func strProps(names ...string) openapi3.Schemas {
	s := openapi3.Schemas{}
	for _, n := range names {
		s[n] = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	}
	return s
}
