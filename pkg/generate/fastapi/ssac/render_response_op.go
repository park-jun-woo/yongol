//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderResponseOp — ResponseOp → Python return 문 렌더링

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
		b.WriteString(fmt.Sprintf("%sreturn %s\n", indent, op.SingleVar))
		return
	}
	b.WriteString(fmt.Sprintf("%sreturn {\n", indent))
	for _, f := range op.Fields {
		b.WriteString(fmt.Sprintf("%s    \"%s\": %s,\n", indent, f.Name, f.Source))
	}
	b.WriteString(fmt.Sprintf("%s}\n", indent))
}
