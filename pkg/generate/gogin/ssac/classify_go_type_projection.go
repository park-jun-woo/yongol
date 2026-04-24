//ff:func feature=gen-gogin type=util control=selection
//ff:what classifyGoTypeProjection — Go 타입 투영 + nullability → pgtypeRowKind

package ssac

import (
	"strings"
)

// classifyGoTypeProjection reverses pkg/parser/ddl/pg_type_to_go.go. The
// parser projected SQL types into a small Go-type alphabet; this function
// reads that projection + nullability and returns the pgx/v5 row-field
// classification. Kept separate from classifyPgtypeRowField so the two
// inference paths (raw SQL type vs. projected Go type) stay testable.
func classifyGoTypeProjection(goType string, notNull bool) pgtypeRowKind {
	switch strings.TrimSpace(goType) {
	case "time.Time":
		return pgTimestamp
	case "json.RawMessage":
		return pgPrimitive
	case "int64", "int32", "int16", "bool", "float64":
		if notNull {
			return pgPrimitive
		}
		return pgUnknown
	case "string":
		// UUID and VARCHAR / TEXT both project to "string" in pgTypeToGo,
		// but only UUID gets a pgtype wrapper in pgx/v5 NOT NULL mode.
		// Without the raw SQL type we conservatively treat NOT NULL
		// string as primitive (VARCHAR / TEXT case). UUID columns are
		// distinguished via the fallback path in pickConvertRHS using
		// apiCastFor(openapi_types.UUID).
		if notNull {
			return pgPrimitive
		}
		return pgTextWrapper
	}
	return pgPrimitive
}
