//ff:func feature=validate type=rule control=sequence topic=ssac-sqlc
//ff:what xqs72CheckEntry — 단일 Input 항목의 OpenAPI int width ↔ sqlc 캐스트 int width 일치 검증

package ssac_sqlc

import (
	"fmt"
	"strings"

	"github.com/ettle/strcase"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xqs72CheckEntry validates a single Input entry for XQS-72.
func xqs72CheckEntry(
	fn ssac.ServiceFunc,
	seq ssac.Sequence,
	key, val string,
	oapiParams map[string]string,
	sqlcParams map[string]bool,
	hasSqlc bool,
	ddlColType map[string]map[string]string,
	tableName string,
	queryBody string,
) (diagnostic.Diagnostic, bool) {
	if !strings.HasPrefix(val, "request.") {
		return diagnostic.Diagnostic{}, false
	}
	paramName := strings.TrimPrefix(val, "request.")

	if hasSqlc && !sqlcParams[key] {
		return diagnostic.Diagnostic{}, false
	}

	colName := strcase.ToSnake(key)
	if _, found := xqs18LookupDDLType(ddlColType, tableName, colName); found {
		return diagnostic.Diagnostic{}, false
	}
	if _, found := xqs18LookupDDLType(ddlColType, tableName, paramName); found {
		return diagnostic.Diagnostic{}, false
	}

	oapiFormat, hasOAPI := oapiParams[paramName]
	if !hasOAPI || (oapiFormat != "int32" && oapiFormat != "int64") {
		return diagnostic.Diagnostic{}, false
	}

	sqlcWidth := inferSqlcParamIntWidth(queryBody, strcase.ToSnake(key))
	if sqlcWidth == "" || oapiFormat == sqlcWidth {
		return diagnostic.Diagnostic{}, false
	}

	return diagnostic.Diagnostic{
		File:  fn.FileName,
		Line:  seq.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XQS-72] Input key %q -- OpenAPI query param format %s != sqlc param inferred type %s",
			key, oapiFormat, sqlcWidth,
		),
		Advice: fmt.Sprintf(
			"Fix OpenAPI to format: %s or add ::%s cast in the sqlc query",
			sqlcWidth, widthToPGCast(oapiFormat),
		),
	}, true
}
