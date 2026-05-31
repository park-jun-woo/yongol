//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what xdo77FS — XDO-77 테스트용 빈 OpenAPILines Fullstack 생성 헬퍼
package openapi_ddl

import (
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdo77FS returns a Fullstack with empty OpenAPI line-index maps for XDO-77 tests.
func xdo77FS() *yongol.Fullstack {
	return &yongol.Fullstack{
		OpenAPILines: &oapiparser.LineIndex{
			SchemaProperties: map[string]map[string]int{},
			Schemas:          map[string]int{},
		},
	}
}
