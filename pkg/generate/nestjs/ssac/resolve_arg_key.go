//ff:func feature=gen-nestjs type=util control=sequence
//ff:what resolveArgKey — FieldArg.ColumnName 우선, fallback toSnake(Key) 로 Prisma where 키 추출

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// resolveArgKey extracts a snake_case key from a FieldArg for Prisma where
// clauses. Prefers ColumnName (DDL-origin snake_case from Phase018 IR),
// then Key, then Field — each converted via toSnake fallback.
func resolveArgKey(a ir.FieldArg) string {
	if a.ColumnName != "" {
		return a.ColumnName
	}
	if a.Key != "" {
		return toSnake(a.Key)
	}
	key := strings.TrimPrefix(a.Field, ".")
	if key == "" {
		return "id"
	}
	return toSnake(key)
}
