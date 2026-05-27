//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderCallOp — CallOp → DI 서비스 메서드 호출 TypeScript await 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderCallOp writes a DI-injected service method call with optional result
// binding. External packages are referenced via this.{pkg}Service.{method}().
func renderCallOp(b *strings.Builder, op *ir.CallOp, indent string) {
	if op == nil {
		return
	}
	args := renderCallArgs(op.Args)
	caller := formatCallTarget(op.Package, op.Function)
	if op.ResultVar != "" {
		b.WriteString(fmt.Sprintf("%sconst %s = await %s(%s);\n",
			indent, op.ResultVar, caller, args))
	} else {
		b.WriteString(fmt.Sprintf("%sawait %s(%s);\n",
			indent, caller, args))
	}
}
