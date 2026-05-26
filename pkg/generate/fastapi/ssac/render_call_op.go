//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderCallOp — CallOp → 외부 함수 호출 Python await 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderCallOp writes an external function call with optional result binding.
func renderCallOp(b *strings.Builder, op *ir.CallOp, indent string) {
	if op == nil {
		return
	}
	pkg := op.Package
	if pkg != "" {
		pkg += "."
	}
	args := renderCallArgs(op.Args)
	if op.ResultVar != "" {
		b.WriteString(fmt.Sprintf("%s%s = await %s%s(%s)\n",
			indent, op.ResultVar, pkg, snakeCase(op.Function), args))
	} else {
		b.WriteString(fmt.Sprintf("%sawait %s%s(%s)\n",
			indent, pkg, snakeCase(op.Function), args))
	}
}
