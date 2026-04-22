//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XDO-9 — components.schemas의 property가 대응 DDL 컬럼에 없으면 ERROR (ghost property)

package openapi_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdo09GhostProperty validates XDO-9: every property of a components.schemas
// entry whose name maps to a DDL table must correspond to an actual DDL column.
func xdo09GhostProperty(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Components == nil || fs.OpenAPIDoc.Components.Schemas == nil {
		return nil
	}

	tables := make(map[string]map[string]bool, len(fs.DDLTables))
	for _, t := range fs.DDLTables {
		cols := make(map[string]bool, len(t.Columns))
		for c := range t.Columns {
			cols[c] = true
		}
		tables[t.Name] = cols
	}

	var diags []diagnostic.Diagnostic
	for schemaName, schemaRef := range fs.OpenAPIDoc.Components.Schemas {
		diags = append(diags, xdo09ScanSchemaProps(fs, schemaName, schemaRef, tables)...)
	}
	return diags
}
