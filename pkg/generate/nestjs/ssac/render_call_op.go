//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderCallOp — CallOp → 외부 함수 호출 TypeScript await 문 렌더링

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
		b.WriteString(fmt.Sprintf("%sconst %s = await %s%s(%s);\n",
			indent, op.ResultVar, pkg, op.Function, args))
	} else {
		b.WriteString(fmt.Sprintf("%sawait %s%s(%s);\n",
			indent, pkg, op.Function, args))
	}
}
