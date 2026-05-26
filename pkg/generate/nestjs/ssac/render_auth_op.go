//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderAuthOp — AuthOp → OPA 정책 평가 TODO 주석 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderAuthOp writes an OPA policy evaluation placeholder comment.
func renderAuthOp(b *strings.Builder, op *ir.AuthOp, indent string) {
	if op == nil {
		return
	}
	b.WriteString(fmt.Sprintf("%s// @auth %s.%s\n", indent, op.Resource, op.Action))
	b.WriteString(fmt.Sprintf("%s// TODO: integrate OPA policy evaluation\n", indent))
}
