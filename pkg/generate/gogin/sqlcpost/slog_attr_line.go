//ff:func feature=gen-gogin type=util control=selection
//ff:what slogAttrLine — sqlc Go 타입별 slog.<Type> 생성 코드 라인 선택

package sqlcpost

import "fmt"

// slogAttrLine picks the slog.<Type> constructor that matches the sqlc Go type.
// Unknown types fall back to slog.Any for safety.
func slogAttrLine(colName, fieldName, goType string) string {
	switch goType {
	case "int64":
		return fmt.Sprintf("slog.Int64(%q, r.%s)", colName, fieldName)
	case "string":
		return fmt.Sprintf("slog.String(%q, r.%s)", colName, fieldName)
	case "bool":
		return fmt.Sprintf("slog.Bool(%q, r.%s)", colName, fieldName)
	case "time.Time":
		return fmt.Sprintf("slog.Time(%q, r.%s)", colName, fieldName)
	case "float64":
		return fmt.Sprintf("slog.Float64(%q, r.%s)", colName, fieldName)
	default:
		return fmt.Sprintf("slog.Any(%q, r.%s)", colName, fieldName)
	}
}
