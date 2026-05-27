//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderAuthOp — AuthOp → authz_check 호출 Python 코드 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderAuthOp writes an authz_check call for OPA policy evaluation.
func renderAuthOp(b *strings.Builder, op *ir.AuthOp, indent string) {
	if op == nil {
		return
	}
	b.WriteString(fmt.Sprintf("%sawait authz_check(\n", indent))
	b.WriteString(fmt.Sprintf("%s    user,\n", indent))
	b.WriteString(fmt.Sprintf("%s    action=\"%s\",\n", indent, op.Action))
	b.WriteString(fmt.Sprintf("%s    resource=\"%s\",\n", indent, op.Resource))
	for _, input := range op.Inputs {
		b.WriteString(fmt.Sprintf("%s    %s=%s,\n", indent, snakeCase(input.Key), renderArgValue(input)))
	}
	b.WriteString(fmt.Sprintf("%s)\n", indent))
}
