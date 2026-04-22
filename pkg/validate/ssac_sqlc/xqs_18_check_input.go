//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what xqs18CheckInput — Inputs map 항목 (request.<param>) 의 OpenAPI↔DDL 타입 호환성 판정

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
			"[XQS-18] Input key %q 의 OpenAPI 타입 %q과 sqlc/DDL 타입 %q이 불일치합니다",
			key, oapiType, goType,
		),
		Advice: "OpenAPI param 타입을 맞추거나 DDL 컬럼 타입을 변경하세요",
	}, true
}
