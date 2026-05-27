//ff:func feature=gen-fastapi type=util control=sequence
//ff:what resolveDataKey — FieldArg → SQLAlchemy data 절 snake_case 키 추출

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// resolveDataKey extracts a snake_case key for the data clause from a FieldArg.
func resolveDataKey(a ir.FieldArg) string {
	if a.Key != "" {
		return snakeCase(a.Key)
	}
	raw := strings.TrimPrefix(a.Field, ".")
	if raw == "" {
		return ""
	}
	return snakeCase(raw)
}
