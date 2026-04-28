//ff:func feature=manifest type=util control=selection
//ff:what pgTypeToGo — PostgreSQL 타입을 Go 타입으로 매핑
package ddl

import "strings"

// pgTypeToGo maps a raw, uppercased PostgreSQL type token to the Go type the
// parser used to project columns to. Parser no longer calls this — it's
// kept for the GoTypeOf wrapper and for compat consumers that want the
// projection from a raw token.
func pgTypeToGo(pgType string) string {
	switch pgType {
	case "BIGINT", "BIGSERIAL", "INTEGER", "SERIAL", "INT", "SMALLINT":
		return "int64"
	case "VARCHAR", "TEXT", "UUID", "CHAR":
		return "string"
	case "BOOLEAN", "BOOL":
		return "bool"
	case "TIMESTAMPTZ", "TIMESTAMP", "DATE":
		return "time.Time"
	case "NUMERIC", "DECIMAL", "REAL", "FLOAT", "DOUBLE":
		return "float64"
	case "JSONB", "JSON":
		return "json.RawMessage"
	default:
		if strings.HasPrefix(pgType, "VARCHAR") || strings.HasPrefix(pgType, "CHAR") {
			return "string"
		}
		return "string"
	}
}
