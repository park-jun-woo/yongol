//ff:func feature=migration type=parser control=selection
//ff:what NormalizeType — DDL 타입 문자열을 CanonicalType 으로 변환 (aliases·SERIAL 해체)
package migration

import (
	"strings"
)

// NormalizeType converts a raw DDL type string (e.g. "varchar(255)",
// "integer", "TIMESTAMP WITH TIME ZONE", "int4", "BIGSERIAL") into a
// CanonicalType. Unknown bases fall through with their uppercase literal.
//
// Returns (CanonicalType, isSerial).  isSerial is true for SERIAL /
// BIGSERIAL / SMALLSERIAL so callers can attach a nextval() default.
func NormalizeType(raw string) (CanonicalType, bool) {
	s := strings.TrimSpace(raw)
	if s == "" {
		return CanonicalType{}, false
	}

	// Array suffix `[]` (possibly with spaces).
	array := false
	for strings.HasSuffix(s, "[]") {
		array = true
		s = strings.TrimSpace(strings.TrimSuffix(s, "[]"))
	}

	// Split off parameter list e.g. VARCHAR(255), NUMERIC(10,2).
	base := s
	var params string
	if i := strings.Index(s, "("); i >= 0 && strings.HasSuffix(s, ")") {
		base = strings.TrimSpace(s[:i])
		params = s[i+1 : len(s)-1]
	}

	upper := strings.ToUpper(strings.Join(strings.Fields(base), " "))

	ct := CanonicalType{Array: array}
	isSerial := false

	switch upper {
	case "INT", "INT4", "INTEGER":
		ct.Base = "INTEGER"
	case "BIGINT", "INT8":
		ct.Base = "BIGINT"
	case "SMALLINT", "INT2":
		ct.Base = "SMALLINT"
	case "BOOL", "BOOLEAN":
		ct.Base = "BOOLEAN"
	case "VARCHAR", "CHARACTER VARYING":
		ct.Base = "VARCHAR"
	case "CHAR", "CHARACTER":
		ct.Base = "CHAR"
	case "TEXT":
		ct.Base = "TEXT"
	case "UUID":
		ct.Base = "UUID"
	case "JSONB":
		ct.Base = "JSONB"
	case "JSON":
		ct.Base = "JSON"
	case "BYTEA":
		ct.Base = "BYTEA"
	case "NUMERIC", "DECIMAL":
		ct.Base = "NUMERIC"
	case "REAL", "FLOAT4":
		ct.Base = "REAL"
	case "DOUBLE PRECISION", "FLOAT8":
		ct.Base = "DOUBLE PRECISION"
	case "DATE":
		ct.Base = "DATE"
	case "TIME", "TIME WITHOUT TIME ZONE":
		ct.Base = "TIME"
	case "TIMETZ", "TIME WITH TIME ZONE":
		ct.Base = "TIMETZ"
	case "TIMESTAMP", "TIMESTAMP WITHOUT TIME ZONE":
		ct.Base = "TIMESTAMP"
	case "TIMESTAMPTZ", "TIMESTAMP WITH TIME ZONE":
		ct.Base = "TIMESTAMPTZ"
	case "SERIAL", "SERIAL4":
		ct.Base = "INTEGER"
		isSerial = true
	case "BIGSERIAL", "SERIAL8":
		ct.Base = "BIGINT"
		isSerial = true
	case "SMALLSERIAL", "SERIAL2":
		ct.Base = "SMALLINT"
		isSerial = true
	default:
		ct.Base = upper
	}

	if params != "" {
		switch ct.Base {
		case "VARCHAR", "CHAR":
			ct.Length = parseIntSafe(strings.TrimSpace(params))
		case "NUMERIC":
			parts := strings.Split(params, ",")
			ct.Precision = parseIntSafe(strings.TrimSpace(parts[0]))
			if len(parts) > 1 {
				ct.Scale = parseIntSafe(strings.TrimSpace(parts[1]))
			}
		}
	}

	return ct, isSerial
}

func parseIntSafe(s string) int {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		n = n*10 + int(r-'0')
	}
	return n
}
