//ff:func feature=gen-nestjs type=util control=selection
//ff:what mapPGFamily — 정규화된 PG 타입명 → Prisma 스칼라 타입 매핑

package prisma

// mapPGFamily maps a normalized PG type name to a Prisma scalar type.
func mapPGFamily(upper string) string {
	switch upper {
	case "BIGINT", "BIGSERIAL":
		return "BigInt"
	case "INTEGER", "INT", "SERIAL", "INT4", "SMALLINT", "INT2":
		return "Int"
	case "TEXT", "VARCHAR", "CHAR", "CHARACTER VARYING", "CITEXT":
		return "String"
	case "BOOLEAN", "BOOL":
		return "Boolean"
	case "TIMESTAMP", "TIMESTAMP WITHOUT TIME ZONE", "TIMESTAMPTZ",
		"TIMESTAMP WITH TIME ZONE", "DATE":
		return "DateTime"
	case "UUID":
		return "String"
	case "JSONB", "JSON":
		return "Json"
	case "NUMERIC", "DECIMAL":
		return "Decimal"
	case "FLOAT", "DOUBLE PRECISION", "REAL", "FLOAT4", "FLOAT8":
		return "Float"
	case "BYTEA":
		return "Bytes"
	case "INET", "INTERVAL":
		return "String"
	default:
		return "String"
	}
}
