//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what renderPrismaData — FieldArg 배열 → Prisma data 절 문자열 생성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderPrismaData builds a Prisma data clause from FieldArgs.
func renderPrismaData(args []ir.FieldArg) string {
	if len(args) == 0 {
		return "{ data: body }"
	}
	parts := make([]string, 0, len(args))
	for _, a := range args {
		key := resolveDataKey(a)
		if key == "" {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s: %s", key, renderArgValue(a)))
	}
	if len(parts) == 0 {
		return "{ data: body }"
	}
	return "{ data: { " + strings.Join(parts, ", ") + " } }"
}
