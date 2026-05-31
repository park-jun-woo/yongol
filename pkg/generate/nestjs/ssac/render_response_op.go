//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what renderResponseOp — ResponseOp → return 문 렌더링 (tsSourceExpr 으로 PascalCase 필드를 camelCase 변환)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderResponseOp writes a return statement for the response.
func renderResponseOp(b *strings.Builder, op *ir.ResponseOp, indent string) {
	if op == nil {
		return
	}
	if op.SingleVar != "" {
		b.WriteString(fmt.Sprintf("%sreturn %s;\n", indent, op.SingleVar))
		return
	}
	b.WriteString(fmt.Sprintf("%sreturn {\n", indent))
	for _, f := range op.Fields {
		b.WriteString(fmt.Sprintf("%s  %s: %s,\n", indent, f.Name, tsSourceExpr(f.Source)))
	}
	b.WriteString(fmt.Sprintf("%s};\n", indent))
}
