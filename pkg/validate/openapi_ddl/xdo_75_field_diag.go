//ff:func feature=validate type=util control=sequence topic=openapi-ddl
//ff:what xdo75FieldDiag — determines whether a single field violates XDO-75 and generates the diagnostic

package openapi_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// xdo75FieldDiag performs the per-field check. Returns (diag, true) when the
// field violates XDO-75 (optional in OpenAPI, NOT NULL in DDL, no DEFAULT).
func xdo75FieldDiag(fs *yongol.Fullstack, opID, fieldName string, fc oapiparser.FieldConstraint) (diagnostic.Diagnostic, bool) {
	if fc.Required {
		// required fields are handled by XDO-76, skip here.
		return diagnostic.Diagnostic{}, false
	}
	col := caseconv.PascalToSnake(fieldName)
	tbl := findDDLTableWithColumn(fs, col)
	if tbl == nil {
		return diagnostic.Diagnostic{}, false
	}
	if !isColumnNotNull(tbl, col) {
		return diagnostic.Diagnostic{}, false
	}
	if c, ok := tbl.Columns[col]; ok && c.HasDefault {
		return diagnostic.Diagnostic{}, false
	}
	line := fs.OpenAPILines.RequestFieldLine(opID, fieldName)
	if line == 0 {
		line = fc.Line
	}
	return diagnostic.Diagnostic{
		File:  "api/openapi.yaml",
		Line:  line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XDO-75] %s — OpenAPI optional field %q is NOT NULL in DDL table %q with no DEFAULT",
			opID, fieldName, tbl.Name,
		),
		Advice: "Add a DEFAULT to the DDL column, or change the field to required in OpenAPI",
	}, true
}
