//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what renderPrismaWhere — FieldArg 배열 → Prisma where 절 문자열 생성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderPrismaWhere builds a Prisma where clause from FieldArgs.
func renderPrismaWhere(args []ir.FieldArg) string {
	if len(args) == 0 {
		return "{}"
	}
	parts := make([]string, 0, len(args))
	for _, a := range args {
		key := resolveArgKey(a)
		parts = append(parts, fmt.Sprintf("%s: %s", key, renderArgValue(a)))
	}
	return "{ where: { " + strings.Join(parts, ", ") + " } }"
}
