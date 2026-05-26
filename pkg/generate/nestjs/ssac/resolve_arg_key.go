//ff:func feature=gen-nestjs type=util control=sequence
//ff:what resolveArgKey — FieldArg → Prisma where 절 키 추출

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// resolveArgKey extracts a key from a FieldArg, preferring Key, then Field.
func resolveArgKey(a ir.FieldArg) string {
	if a.Key != "" {
		return a.Key
	}
	key := strings.TrimPrefix(a.Field, ".")
	if key == "" {
		return "id"
	}
	return key
}
