//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderEvalOp — EvalOp → DI 서비스 메서드 bool 평가 if 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderEvalOp writes a boolean function evaluation guard using DI-injected
// services.
func renderEvalOp(b *strings.Builder, op *ir.EvalOp, indent string) {
	if op == nil {
		return
	}
	args := renderCallArgs(op.Args)
	caller := formatCallTarget(op.Package, op.Function)
	b.WriteString(fmt.Sprintf("%sif (await %s(%s)) {\n",
		indent, caller, args))
	b.WriteString(fmt.Sprintf("%s  throw new HttpException('%s', HttpStatus.%s);\n",
		indent, op.Message, httpStatusConst(op.StatusCode)))
	b.WriteString(fmt.Sprintf("%s}\n", indent))
}
