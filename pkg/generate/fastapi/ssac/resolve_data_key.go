//ff:func feature=gen-fastapi type=util control=sequence
//ff:what resolveDataKey — FieldArg.ColumnName 우선, fallback snakeCase(Key) 로 SQLAlchemy data 키 추출

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// resolveDataKey extracts a snake_case key for the data clause from a FieldArg.
// Prefers ColumnName (DDL-origin snake_case from Phase018 IR), then Key, then
// Field.
func resolveDataKey(a ir.FieldArg) string {
	if a.ColumnName != "" {
		return a.ColumnName
	}
	if a.Key != "" {
		return snakeCase(a.Key)
	}
	raw := strings.TrimPrefix(a.Field, ".")
	if raw == "" {
		return ""
	}
	return snakeCase(raw)
}
