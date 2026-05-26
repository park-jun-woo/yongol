//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderEmptyOp — EmptyOp → null 가드 TypeScript if 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderEmptyOp writes a null guard that throws HttpException.
func renderEmptyOp(b *strings.Builder, op *ir.EmptyOp, indent string) {
	if op == nil {
		return
	}
	b.WriteString(fmt.Sprintf("%sif (!%s) {\n", indent, op.VarName))
	b.WriteString(fmt.Sprintf("%s  throw new HttpException('%s', HttpStatus.%s);\n",
		indent, op.Message, httpStatusConst(op.StatusCode)))
	b.WriteString(fmt.Sprintf("%s}\n", indent))
}
