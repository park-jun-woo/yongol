//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderPostOp — PostOp → Prisma create TypeScript 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderPostOp writes a Prisma create call.
func renderPostOp(b *strings.Builder, op *ir.PostOp, indent, prismaRef string) {
	if op == nil {
		return
	}
	model := lcFirst(op.Model)
	data := renderPrismaData(op.Args)
	b.WriteString(fmt.Sprintf("%sconst %s = await %s.%s.create(%s);\n",
		indent, op.VarName, prismaRef, model, data))
}
