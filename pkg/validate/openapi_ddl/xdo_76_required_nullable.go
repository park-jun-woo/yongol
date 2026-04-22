//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XDO-76 — OpenAPI required + DDL nullable → WARNING. `-- @nullable` 어노테이션 면제.

package openapi_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// xdo76RequiredNullable validates XDO-76: if an OpenAPI request field is
// required but the corresponding DDL column allows NULL (no NOT NULL, not PK),
// it may indicate a schema mismatch. Columns annotated with `-- @nullable`
// are exempt as intentional design.
func xdo76RequiredNullable(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic

	for opID, fields := range fs.RequestConstraints {
		for fieldName, fc := range fields {
			if !fc.Required {
				continue
			}
			col := caseconv.PascalToSnake(fieldName)

			tbl := findDDLTableWithColumn(fs, col)
			if tbl == nil {
				continue
			}

			// If the column is NOT NULL (or PK), it's fine.
			if isColumnNotNull(tbl, col) {
				continue
			}

			// Exempt columns annotated with -- @nullable.
			if tbl.NullableAnnot != nil && tbl.NullableAnnot[col] {
				continue
			}

			line := fs.OpenAPILines.RequestFieldLine(opID, fieldName)
			if line == 0 {
				line = fc.Line
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:  "api/openapi.yaml",
				Line:  line,
				Phase: diagnostic.PhaseValidate,
				Level: diagnostic.LevelWarning,
				Message: fmt.Sprintf(
					"[XDO-76] %s — OpenAPI required field %q is nullable in DDL table %q",
					opID, fieldName, tbl.Name,
				),
				Advice: "DDL에 NOT NULL을 추가하거나 의도적이면 -- @nullable 어노테이션을 추가하세요",
			})
		}
	}
	return diags
}
