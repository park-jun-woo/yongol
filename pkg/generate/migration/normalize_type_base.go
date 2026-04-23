//ff:func feature=migration type=parser control=selection
//ff:what normalizeTypeBase — upper base 문자열을 CanonicalType.Base 로 매핑 + SERIAL 여부 반환
package migration

// normalizeTypeBase maps the uppercase base token to its canonical form
// on ct.Base and returns true when the original was SERIAL / BIGSERIAL /
// SMALLSERIAL.
func normalizeTypeBase(upper string, ct *CanonicalType) bool {
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
		return true
	case "BIGSERIAL", "SERIAL8":
		ct.Base = "BIGINT"
		return true
	case "SMALLSERIAL", "SERIAL2":
		ct.Base = "SMALLINT"
		return true
	default:
		ct.Base = upper
	}
	return false
}
