//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what xqs18CheckArg — check OpenAPI↔DDL type compatibility for a single CRUD arg (source=request)

package ssac_sqlc

import (
	"fmt"

	"github.com/ettle/strcase"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xqs18CheckArg validates a single CRUD Arg against the OpenAPI param type
// and the DDL column Go type. Returns (diag, true) on mismatch.
func xqs18CheckArg(fn ssac.ServiceFunc, seq ssac.Sequence, arg ssac.Arg, oapiParams map[string]string, sqlcParams map[string]bool, hasSqlc bool, ddlColType map[string]map[string]string, tableName string) (diagnostic.Diagnostic, bool) {
	if arg.Source != "request" || arg.Field == "" {
		return diagnostic.Diagnostic{}, false
	}
	oapiType, hasOAPI := oapiParams[strcase.ToSnake(arg.Field)]
	if !hasOAPI {
		oapiType, hasOAPI = oapiParams[arg.Field]
	}
	if !hasOAPI {
		return diagnostic.Diagnostic{}, false
	}
	compatible, hasCompat := openAPITypeCompatible[oapiType]
	if !hasCompat {
		return diagnostic.Diagnostic{}, false
	}
	// sqlc param field name is arg.Field (PascalCase)
	if hasSqlc && !sqlcParams[arg.Field] {
		return diagnostic.Diagnostic{}, false
	}
	colName := strcase.ToSnake(arg.Field)
	goType, found := xqs18LookupDDLType(ddlColType, tableName, colName)
	if !found {
		return diagnostic.Diagnostic{}, false
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
			arg.Field, oapiType, goType,
		),
		Advice: "Align the OpenAPI param type with the DDL column type, or change the column type",
	}, true
}
