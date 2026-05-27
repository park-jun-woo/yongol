//ff:func feature=gen-fastapi type=util control=sequence
//ff:what resolveArgKey — FieldArg.ColumnName 우선, fallback snakeCase(Key) 로 SQLAlchemy where 키 추출

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// resolveArgKey extracts a snake_case key from a FieldArg for SQLAlchemy
// where clauses. Prefers ColumnName (DDL-origin snake_case from Phase018 IR),
// then Key, then Field — each converted via snakeCase fallback.
func resolveArgKey(a ir.FieldArg) string {
	if a.ColumnName != "" {
		return a.ColumnName
	}
	if a.Key != "" {
		return snakeCase(a.Key)
	}
	key := strings.TrimPrefix(a.Field, ".")
	if key == "" {
		return "id"
	}
	return snakeCase(key)
}
