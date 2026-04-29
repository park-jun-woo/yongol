//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what modelToTableName — 모델 이름(예: "User") → DDL 테이블 이름(예: "users")

package ssac_sqlc

import (
	"github.com/ettle/strcase"
	"github.com/jinzhu/inflection"
)

// modelToTableName converts a SSaC / sqlc Model identifier to the canonical
// DDL table name. Mirrors the convention used in pkg/validate/ssac_ddl
// (PascalCase singular → snake_case plural). Examples:
//
//	"User"        → "users"
//	"AuditLog"    → "audit_logs"
//	"Workflow"    → "workflows"
func modelToTableName(model string) string {
	return inflection.Plural(strcase.ToSnake(model))
}
