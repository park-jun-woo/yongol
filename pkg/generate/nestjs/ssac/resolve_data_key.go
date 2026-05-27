//ff:func feature=gen-nestjs type=util control=sequence
//ff:what resolveDataKey — FieldArg.ColumnName 우선, fallback toSnake(Key) 로 Prisma data 키 추출

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// resolveDataKey extracts a snake_case key for the Prisma data clause from a
// FieldArg. Prefers ColumnName (DDL-origin snake_case from Phase018 IR),
// then Key, then Field.
func resolveDataKey(a ir.FieldArg) string {
	if a.ColumnName != "" {
		return a.ColumnName
	}
	if a.Key != "" {
		return toSnake(a.Key)
	}
	raw := strings.TrimPrefix(a.Field, ".")
	if raw == "" {
		return ""
	}
	return toSnake(raw)
}
