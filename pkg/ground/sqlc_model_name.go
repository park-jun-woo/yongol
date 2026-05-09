//ff:func feature=ground type=util control=sequence topic=ddl
//ff:what sqlcModelName — 복수형 snake_case 테이블명 → sqlc PascalCase 단수 모델명 변환

package ground

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// sqlcModelName converts a plural snake_case table name to its sqlc model
// PascalCase singular form. Example: "workflows" → "Workflow",
// "execution_logs" → "ExecutionLog", "organizations" → "Organization".
func sqlcModelName(tableName string) string {
	singular := singularize(tableName)
	return caseconv.SnakeToPascalSqlc(singular)
}

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
