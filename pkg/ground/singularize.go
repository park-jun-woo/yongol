//ff:func feature=ground type=util control=selection topic=ddl
//ff:what singularize — 영어 복수형 snake_case 테이블명에서 단수형 접미사 제거

package ground

import "strings"

// singularize removes the English plural suffix from a lower-snake table
// name. Matches the same rules as ddlTableSingular in
// pkg/generate/gogin/ssac/ddl_table_singular.go.
func singularize(name string) string {
	switch {
	case strings.HasSuffix(name, "ies"):
		return name[:len(name)-3] + "y"
	case strings.HasSuffix(name, "sses"):
		return name[:len(name)-2]
	case strings.HasSuffix(name, "xes"):
		return name[:len(name)-2]
	case strings.HasSuffix(name, "s") && !strings.HasSuffix(name, "ss"):
		return name[:len(name)-1]
	default:
		return name
	}
}
