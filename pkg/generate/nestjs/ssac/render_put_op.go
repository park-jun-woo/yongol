//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderPutOp — PutOp → Prisma update TypeScript 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderPutOp writes a Prisma update call.
func renderPutOp(b *strings.Builder, op *ir.PutOp, indent, prismaRef string) {
	if op == nil {
		return
	}
	model := lcFirst(op.Model)
	b.WriteString(fmt.Sprintf("%sawait %s.%s.update({ where: { id: params.id }, data: body });\n",
		indent, prismaRef, model))
}
