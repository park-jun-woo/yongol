//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderVerifyPasswordOp — VerifyPasswordOp → Prisma lookup + bcrypt 검증 TypeScript 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderVerifyPasswordOp writes a user lookup and password verification block.
func renderVerifyPasswordOp(b *strings.Builder, op *ir.VerifyPasswordOp, indent, prismaRef string) {
	if op == nil {
		return
	}
	model := lcFirst(op.Model)
	b.WriteString(fmt.Sprintf("%sconst %s = await %s.%s.findUnique({ where: { %s: %s } });\n",
		indent, op.ResultVar, prismaRef, model, lcFirst(op.EmailCol), op.EmailExpr))
	b.WriteString(fmt.Sprintf("%sif (!%s) {\n", indent, op.ResultVar))
	b.WriteString(fmt.Sprintf("%s  throw new HttpException('%s', HttpStatus.UNAUTHORIZED);\n",
		indent, op.Message))
	b.WriteString(fmt.Sprintf("%s}\n", indent))
	b.WriteString(fmt.Sprintf("%s// TODO: bcrypt.compare(%s, %s.%s)\n",
		indent, op.PasswordExpr, op.ResultVar, lcFirst(op.HashCol)))
}
