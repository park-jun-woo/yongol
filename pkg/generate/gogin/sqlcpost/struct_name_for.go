//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what structNameFor — DDL 테이블명을 sqlc 생성 Go struct 이름으로 매핑 (단수화 + PascalCase)

package sqlcpost

import "strings"

// structNameFor maps a DDL table name to the sqlc-generated Go struct name.
// sqlc singularizes the table name then PascalCases it on underscore.
// E.g. users → User, audit_logs → AuditLog, execution_logs → ExecutionLog.
func structNameFor(table string) string {
	singular := singularize(table)
	parts := strings.Split(singular, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]) + p[1:])
	}
	return b.String()
}
