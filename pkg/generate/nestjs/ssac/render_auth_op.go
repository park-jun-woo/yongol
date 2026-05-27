//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what renderAuthOp — AuthOp → AuthzService.check 호출 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderAuthOp writes an authz service check call.
func renderAuthOp(b *strings.Builder, op *ir.AuthOp, indent string) {
	if op == nil {
		return
	}
	b.WriteString(fmt.Sprintf("%sawait this.authz.check({\n", indent))
	b.WriteString(fmt.Sprintf("%s  action: '%s',\n", indent, op.Action))
	b.WriteString(fmt.Sprintf("%s  resource: '%s',\n", indent, op.Resource))
	for _, input := range op.Inputs {
		b.WriteString(fmt.Sprintf("%s  %s: %s,\n", indent, input.Key, renderArgValue(input)))
	}
	b.WriteString(fmt.Sprintf("%s});\n", indent))
}
