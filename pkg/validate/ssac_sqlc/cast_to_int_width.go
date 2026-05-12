//ff:func feature=validate type=util control=selection topic=ssac-sqlc
//ff:what castToIntWidth — PostgreSQL 타입 캐스트 토큰을 int32/int64 로 매핑

package ssac_sqlc

// castToIntWidth maps a PostgreSQL type-cast token to "int32" or "int64".
// Empty token (no cast) defaults to "int32".
func castToIntWidth(cast string) string {
	switch cast {
	case "bigint", "int8":
		return "int64"
	case "int", "int4", "integer", "":
		return "int32"
	default:
		// Non-integer cast (e.g. ::text) -- not an int param.
		return ""
	}
}
