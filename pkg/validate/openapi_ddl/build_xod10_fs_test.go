//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what buildXod10FS — xod10 테스트용 Fullstack 빌더 헬퍼

package openapi_ddl

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func buildXod10FS(tables []ddl.Table, schemas openapi3.Schemas) *yongol.Fullstack {
	return &yongol.Fullstack{
		DDLTables: tables,
		OpenAPIDoc: &openapi3.T{
			Components: &openapi3.Components{Schemas: schemas},
		},
		OpenAPILines: &oapiparser.LineIndex{
			SchemaProperties: map[string]map[string]int{},
			Schemas:          map[string]int{},
		},
	}
}
