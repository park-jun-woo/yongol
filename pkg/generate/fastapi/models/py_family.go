//ff:func feature=gen-fastapi type=util control=selection
//ff:what pyFamily — 정규화된 PG 타입명 → Python 타입 매핑

package models

// pyFamily maps a normalized PG type name to a Python type.
func pyFamily(upper string) string {
	switch upper {
	case "BIGINT", "BIGSERIAL", "INTEGER", "INT", "SERIAL", "INT4", "SMALLINT", "INT2":
		return "int"
	case "TEXT", "VARCHAR", "CHAR", "CHARACTER VARYING", "CITEXT":
		return "str"
	case "BOOLEAN", "BOOL":
		return "bool"
	case "TIMESTAMP", "TIMESTAMP WITHOUT TIME ZONE", "TIMESTAMPTZ",
		"TIMESTAMP WITH TIME ZONE":
		return "datetime"
	case "DATE":
		return "date"
	case "UUID":
		return "uuid.UUID"
	case "JSONB", "JSON":
		return "dict[str, Any]"
	case "NUMERIC", "DECIMAL":
		return "Decimal"
	case "FLOAT", "DOUBLE PRECISION", "REAL", "FLOAT4", "FLOAT8":
		return "float"
	case "BYTEA":
		return "bytes"
	case "INET", "CIDR":
		return "str"
	case "INTERVAL":
		return "timedelta"
	default:
		return "str"
	}
}
