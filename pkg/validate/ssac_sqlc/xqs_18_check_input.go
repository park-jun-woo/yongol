//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what xqs18CheckInput — check OpenAPI↔DDL type compatibility for an Inputs map entry (request.<param>)

package ssac_sqlc

import (
	"fmt"
	"strings"

	"github.com/ettle/strcase"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xqs18CheckInput validates a single Inputs map entry where the value is
// prefixed with "request.". Returns (diag, true) on type incompatibility.
func xqs18CheckInput(fn ssac.ServiceFunc, seq ssac.Sequence, key, val string, oapiParams map[string]string, sqlcParams map[string]bool, hasSqlc bool, ddlColType map[string]map[string]string, tableName string) (diagnostic.Diagnostic, bool) {
	if !strings.HasPrefix(val, "request.") {
		return diagnostic.Diagnostic{}, false
	}
	paramName := strings.TrimPrefix(val, "request.")
	oapiType, hasOAPI := oapiParams[paramName]
	if !hasOAPI {
		return diagnostic.Diagnostic{}, false
	}
	compatible, hasCompat := openAPITypeCompatible[oapiType]
	if !hasCompat {
		return diagnostic.Diagnostic{}, false
	}
	// key is the sqlc Params field name (PascalCase)
	if hasSqlc && !sqlcParams[key] {
		return diagnostic.Diagnostic{}, false
	}
	colName := strcase.ToSnake(key)
	goType, found := xqs18LookupDDLType(ddlColType, tableName, colName)
	if !found {
		goType, found = xqs18LookupDDLType(ddlColType, tableName, paramName)
		if !found {
			return diagnostic.Diagnostic{}, false
		}
	}
	if compatible[goType] {
		return diagnostic.Diagnostic{}, false
	}
	return diagnostic.Diagnostic{
		File:  fn.FileName,
		Line:  seq.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XQS-18] Input key %q — OpenAPI type %q does not match sqlc/DDL type %q",
			key, oapiType, goType,
		),
		Advice: "Align the OpenAPI param type with the DDL column type, or change the column type",
	}, true
}
