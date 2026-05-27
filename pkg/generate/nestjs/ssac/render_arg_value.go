//ff:func feature=gen-nestjs type=util control=selection
//ff:what renderArgValue — FieldArg.Location 기반 TypeScript 식별자/리터럴 표현식 생성

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderArgValue produces a TypeScript expression for a FieldArg by
// switching on the enriched Location field from Phase018 IR. When
// Location is empty (OpenAPI absent), falls back to legacy source-based
// mapping via renderArgValueLegacy.
func renderArgValue(a ir.FieldArg) string {
	if a.Literal != "" {
		if a.IsQuoted {
			return fmt.Sprintf("'%s'", a.Literal)
		}
		return a.Literal
	}

	col := a.ColumnName
	if col == "" {
		col = toSnake(fieldName(a))
	}

	switch a.Location {
	case ir.LocPath:
		return fmt.Sprintf("params.%s", col)
	case ir.LocBody:
		return fmt.Sprintf("body.%s", col)
	case ir.LocQuery:
		return fmt.Sprintf("query.%s", col)
	case ir.LocUser:
		return fmt.Sprintf("user.%s", col)
	case ir.LocVar:
		if col != "" && a.Source != "" {
			return fmt.Sprintf("%s.%s", a.Source, col)
		}
		if a.Source != "" {
			return a.Source
		}
		return col
	case ir.LocLiteral:
		return fmt.Sprintf("'%s'", a.Literal)
	}

	// Fallback: Location not set (no OpenAPI doc). Use source-based mapping.
	return renderArgValueLegacy(a, col)
}

// renderArgValueLegacy maps SSaC source names when Location is empty.
func renderArgValueLegacy(a ir.FieldArg, col string) string {
	source := a.Source
	if source == "" {
		source = "params"
	}
	if col == "" {
		switch source {
		case "request":
			return "params"
		case "currentUser":
			return "user"
		default:
			return source
		}
	}
	switch source {
	case "request":
		return fmt.Sprintf("params.%s", col)
	case "currentUser":
		return fmt.Sprintf("user.%s", col)
	default:
		return fmt.Sprintf("%s.%s", source, col)
	}
}

// fieldName extracts the field accessor name from a FieldArg, stripping
// the leading dot.
func fieldName(a ir.FieldArg) string {
	if len(a.Field) > 0 && a.Field[0] == '.' {
		return a.Field[1:]
	}
	return a.Field
}
