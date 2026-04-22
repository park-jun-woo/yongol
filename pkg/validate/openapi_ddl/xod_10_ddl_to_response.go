//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XOD-10 — WARNING when a DDL column is not exposed as a property in the corresponding OpenAPI components.schemas

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
				Advice:  fmt.Sprintf("Add DDL column %q to OpenAPI response schema %s, or annotate the DDL column with -- @sensitive", col, schemaName),
			})
		}
	}
	return diags
}
