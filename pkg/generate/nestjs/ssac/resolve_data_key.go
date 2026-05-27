//ff:func feature=gen-nestjs type=util control=sequence
//ff:what resolveDataKey — FieldArg → Prisma data 절 snake_case 키 추출

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// resolveDataKey extracts a snake_case key for the Prisma data clause from a
// FieldArg.
func resolveDataKey(a ir.FieldArg) string {
	if a.Key != "" {
		return toSnake(a.Key)
	}
	raw := strings.TrimPrefix(a.Field, ".")
	if raw == "" {
		return ""
	}
	return toSnake(raw)
}
