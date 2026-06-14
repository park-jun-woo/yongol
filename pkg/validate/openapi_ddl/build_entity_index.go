//ff:func feature=validate type=loader control=iteration dimension=1 topic=openapi-ddl
//ff:what buildEntityIndex — Fullstack 에서 DDL 테이블/component/SSaC 함수 인덱스 구성

package openapi_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildEntityIndex constructs the lookups used by canonical-response validation.
// schemaForTable maps DDL-table-name → component schema name (XOD-10's mapping),
// reused here as the entity ground truth.
func buildEntityIndex(fs *yongol.Fullstack) *entityIndex {
	idx := &entityIndex{
		g:              fs.Ground(),
		tables:         fs.DDLTables,
		tableExists:    make(map[string]bool, len(fs.DDLTables)),
		schemaForTable: make(map[string]string, len(fs.OpenAPIDoc.Components.Schemas)),
		funcByName:     make(map[string]*ssac.ServiceFunc, len(fs.ServiceFuncs)),
	}
	for _, t := range fs.DDLTables {
		idx.tableExists[t.Name] = true
	}
	for schemaName := range fs.OpenAPIDoc.Components.Schemas {
		idx.schemaForTable[modelToTable(schemaName)] = schemaName
	}
	for i := range fs.ServiceFuncs {
		fn := &fs.ServiceFuncs[i]
		idx.funcByName[fn.Name] = fn
	}
	return idx
}
