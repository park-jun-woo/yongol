//ff:func feature=gen-gogin type=util control=selection
//ff:what GoTypeOf — Column → 단일-토큰 Go 타입 투영 (legacy 호환: int64/string/bool/time.Time/float64/json.RawMessage)

package types

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// GoTypeOf returns the Go-type projection of a parsed Column. It is the
// pre-Phase001 ddl.GoTypeOf helper relocated to the types module so the
// raw → Go projection lives next to the canonical MapPGType matrix.
//
// The projection is intentionally coarse — it collapses every PG family
// onto one of {int64, string, bool, time.Time, float64,
// json.RawMessage} so legacy validators (XDN-04 claim compatibility,
// XDO-77 OpenAPI ↔ DDL field shape, XQS-18 sqlc param ↔ DDL column,
// sqlcpost log emit) can compare against manifest / OpenAPI declared
// Go types without resolving pgtype wrappers. For codegen-side
// decisions use MapPGType + the ConvertExpr / InsertExpr / ResponseExpr
// templates instead — they expose the precise sqlc / api / convert
// shape per family.
func GoTypeOf(col ddl.Column) string {
	t := col.RawType
	if idx := strings.Index(t, "("); idx > 0 {
		t = t[:idx]
	}
	t = strings.TrimSpace(strings.ToUpper(t))
	switch t {
	case "BIGINT", "BIGSERIAL", "INTEGER", "SERIAL", "INT",
		"INT2", "INT4", "INT8", "SMALLINT", "SMALLSERIAL":
		return "int64"
	case "VARCHAR", "TEXT", "UUID", "CHAR", "BPCHAR":
		return "string"
	case "BOOLEAN", "BOOL":
		return "bool"
	case "TIMESTAMPTZ", "TIMESTAMP", "DATE":
		return "time.Time"
	case "NUMERIC", "DECIMAL", "REAL", "FLOAT", "DOUBLE",
		"FLOAT4", "FLOAT8":
		return "float64"
	case "JSONB", "JSON":
		return "json.RawMessage"
	}
	if strings.HasPrefix(t, "VARCHAR") || strings.HasPrefix(t, "CHAR") {
		return "string"
	}
	return "string"
}
