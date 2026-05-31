//ff:func feature=gen-gogin type=test control=sequence
//ff:what lookupHTTPWhat 단위 테스트 (Summary > Description > 기본값 우선순위)
package ssac

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func docWithOp(op *openapi3.Operation) *yongol.Fullstack {
	return &yongol.Fullstack{
		OpenAPIDoc: &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/x", &openapi3.PathItem{Get: op}),
			),
		},
	}
}
