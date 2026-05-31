//ff:func feature=external type=test control=sequence
//ff:what TestExtract* — path/body 파라미터·반환타입·키정렬·오퍼레이션 조회 검증
package external

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func strSchema() *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
}
