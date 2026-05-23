//ff:func feature=gen-gogin type=util control=selection
//ff:what goIntTypeFromSqlcType — pgtype sqlc 타입에 대응하는 Go 정수 타입 반환

package ssac

// goIntTypeFromSqlcType returns the Go integer primitive corresponding to
// a pgtype sqlc type. Returns empty string for non-integer types.
func goIntTypeFromSqlcType(sqlcGoType string) string {
	switch sqlcGoType {
	case "pgtype.Int8":
		return "int64"
	case "pgtype.Int4":
		return "int32"
	case "pgtype.Int2":
		return "int16"
	case "int64":
		return "int64"
	case "int32":
		return "int32"
	default:
		return ""
	}
}
