//ff:func feature=gen-fastapi type=util control=selection
//ff:what renderArgValue — FieldArg.Location 기반 Python 식별자/리터럴 표현식 생성

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderArgValue produces a Python expression for a FieldArg by
// switching on the enriched Location field from Phase018 IR. When
// Location is empty (OpenAPI absent), falls back to legacy source-based
// mapping via renderArgValueLegacy.
func renderArgValue(a ir.FieldArg) string {
	if a.Literal != "" {
		return renderLiteralArg(a)
	}

	col := a.ColumnName
	if col == "" {
		col = snakeCase(fieldName(a))
	}

	switch a.Location {
	case ir.LocPath, ir.LocQuery:
		// FastAPI path/query params are function arguments directly.
		return col
	case ir.LocBody:
		return fmt.Sprintf("body.%s", col)
	case ir.LocUser:
		return fmt.Sprintf("current_user[\"%s\"]", col)
	case ir.LocVar:
		return renderArgValueVar(a, col)
	case ir.LocLiteral:
		return fmt.Sprintf("\"%s\"", a.Literal)
	}

	// Fallback: Location not set (no OpenAPI doc). Use source-based mapping.
	return renderArgValueLegacy(a, col)
}
