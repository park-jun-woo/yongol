//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderDeleteOp — DeleteOp → Prisma delete TypeScript 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderDeleteOp writes a Prisma delete call.
func renderDeleteOp(b *strings.Builder, op *ir.DeleteOp, indent, prismaRef string) {
	if op == nil {
		return
	}
	model := lcFirst(op.Model)
	args := renderPrismaWhere(op.Args)
	b.WriteString(fmt.Sprintf("%sawait %s.%s.delete(%s);\n",
		indent, prismaRef, model, args))
}
