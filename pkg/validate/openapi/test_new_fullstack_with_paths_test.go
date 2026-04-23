//ff:func feature=validate type=test-helper control=sequence topic=openapi-structural
//ff:what 테스트 헬퍼 — openapi3.Paths 만 가진 최소 Fullstack 생성

package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func newFullstackWithPaths(paths *openapi3.Paths) *yongol.Fullstack {
	return &yongol.Fullstack{
		OpenAPIDoc: &openapi3.T{Paths: paths},
		OpenAPILines: &oapiparser.LineIndex{
			Paths: map[string]int{},
		},
	}
}
