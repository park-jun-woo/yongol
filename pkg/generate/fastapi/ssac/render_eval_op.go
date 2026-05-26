//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderEvalOp — EvalOp → bool 함수 평가 if 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderEvalOp writes a boolean function evaluation guard.
func renderEvalOp(b *strings.Builder, op *ir.EvalOp, indent string) {
	if op == nil {
		return
	}
	pkg := op.Package
	if pkg != "" {
		pkg += "."
	}
	args := renderCallArgs(op.Args)
	b.WriteString(fmt.Sprintf("%sif %s%s(%s):\n",
		indent, pkg, snakeCase(op.Function), args))
	b.WriteString(fmt.Sprintf("%s    raise HTTPException(status_code=%d, detail=\"%s\")\n",
		indent, op.StatusCode, op.Message))
}
