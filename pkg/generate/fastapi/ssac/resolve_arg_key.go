//ff:func feature=gen-fastapi type=util control=sequence
//ff:what resolveArgKey — FieldArg → SQLAlchemy where 절 snake_case 키 추출

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// resolveArgKey extracts a snake_case key from a FieldArg for SQLAlchemy
// where clauses. Prefers Key (converted to snake_case), then falls back to
// Field converted to snake_case.
func resolveArgKey(a ir.FieldArg) string {
	if a.Key != "" {
		return snakeCase(a.Key)
	}
	key := strings.TrimPrefix(a.Field, ".")
	if key == "" {
		return "id"
	}
	return snakeCase(key)
}
