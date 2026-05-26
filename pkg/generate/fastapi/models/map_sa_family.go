//ff:func feature=gen-fastapi type=util control=selection
//ff:what mapSAFamily — 정규화된 PG 타입명 → SQLAlchemy 타입 매핑

package models

// mapSAFamily maps a normalized PG type name to a SQLAlchemy type name.
func mapSAFamily(upper string) string {
	switch upper {
	case "BIGINT", "BIGSERIAL":
		return "Integer"
	case "INTEGER", "INT", "SERIAL", "INT4", "SMALLINT", "INT2":
		return "Integer"
	case "TEXT":
		return "Text"
	case "VARCHAR", "CHAR", "CHARACTER VARYING", "CITEXT":
		return "String"
	case "BOOLEAN", "BOOL":
		return "Boolean"
	case "TIMESTAMP", "TIMESTAMP WITHOUT TIME ZONE":
		return "DateTime(timezone=False)"
	case "TIMESTAMPTZ", "TIMESTAMP WITH TIME ZONE":
		return "DateTime(timezone=True)"
	case "DATE":
		return "Date"
	case "UUID":
		return "Uuid"
	case "JSONB", "JSON":
		return "JSONB"
	case "NUMERIC", "DECIMAL":
		return "Numeric"
	case "FLOAT", "DOUBLE PRECISION", "REAL", "FLOAT4", "FLOAT8":
		return "Float"
	case "BYTEA":
		return "LargeBinary"
	case "INET", "CIDR":
		return "INET"
	case "INTERVAL":
		return "Interval"
	default:
		return "String"
	}
}
