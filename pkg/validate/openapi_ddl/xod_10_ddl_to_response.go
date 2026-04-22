//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XOD-10 — DDL 컬럼이 대응 OpenAPI components.schemas property 로 노출되지 않으면 WARNING

package openapi_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xod10DDLToResponse validates XOD-10: every DDL column on a table whose
// pluralised name maps to a components.schemas entry should appear as a
// property of that schema. `-- @sensitive` / `-- @archived` columns are exempt.
func xod10DDLToResponse(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Components == nil || fs.OpenAPIDoc.Components.Schemas == nil {
		return nil
	}
	g := fs.Ground()

	schemaForTable := make(map[string]string, len(fs.OpenAPIDoc.Components.Schemas))
	for schemaName := range fs.OpenAPIDoc.Components.Schemas {
		schemaForTable[modelToTable(schemaName)] = schemaName
	}

	var diags []diagnostic.Diagnostic
	for _, t := range fs.DDLTables {
		schemaName, ok := schemaForTable[t.Name]
		if !ok {
			continue
		}
		schemaRef := fs.OpenAPIDoc.Components.Schemas[schemaName]
		if schemaRef == nil || schemaRef.Value == nil {
			continue
		}
		props := schemaRef.Value.Properties
		for col := range t.Columns {
			if _, exists := props[col]; exists {
				continue
			}
			if g != nil && (g.Flags["sensitive."+t.Name+"."+col] || g.Flags["archived."+t.Name+"."+col]) {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    "api/openapi.yaml",
				Line:    fs.OpenAPILines.SchemaLine(schemaName),
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XOD-10] DDL column %s.%s missing from OpenAPI schema %s", t.Name, col, schemaName),
				Advice:  fmt.Sprintf("DDL 컬럼 %q 를 OpenAPI response schema %s 에 추가하거나 DDL 컬럼에 -- @sensitive 어노테이션을 다세요", col, schemaName),
			})
		}
	}
	return diags
}
