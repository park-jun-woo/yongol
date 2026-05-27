//ff:func feature=gen-nestjs type=util control=sequence
//ff:what resolveArgKey — FieldArg → Prisma where 절 snake_case 키 추출

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// resolveArgKey extracts a snake_case key from a FieldArg for Prisma where
// clauses. Prefers Key (already snake_case from SSaC Inputs map), then falls
// back to Field converted to snake_case.
func resolveArgKey(a ir.FieldArg) string {
	if a.Key != "" {
		return toSnake(a.Key)
	}
	key := strings.TrimPrefix(a.Field, ".")
	if key == "" {
		return "id"
	}
	return toSnake(key)
}
