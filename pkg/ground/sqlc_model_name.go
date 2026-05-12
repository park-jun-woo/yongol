//ff:func feature=ground type=util control=sequence topic=ddl
//ff:what sqlcModelName — 복수형 snake_case 테이블명 → sqlc PascalCase 단수 모델명 변환

package ground

import (
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// sqlcModelName converts a plural snake_case table name to its sqlc model
// PascalCase singular form. Example: "workflows" → "Workflow",
// "execution_logs" → "ExecutionLog", "organizations" → "Organization".
func sqlcModelName(tableName string) string {
	singular := singularize(tableName)
	return caseconv.SnakeToPascalSqlc(singular)
}
