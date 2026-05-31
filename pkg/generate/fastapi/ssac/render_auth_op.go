//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderAuthOp — AuthOp.Ownership 기반 DB lookup + authz_check 호출 Python 코드 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderAuthOp writes an authz_check call for OPA policy evaluation. When
// Ownership is populated (from Phase018 IR Rego enrichment), an ownership
// DB lookup query is emitted before the authz check.
func renderAuthOp(b *strings.Builder, op *ir.AuthOp, indent string) {
	if op == nil {
		return
	}

	if op.Ownership != nil {
		writeOwnershipLookup(b, op.Ownership, indent)
	}

	b.WriteString(fmt.Sprintf("%sawait authz_check(\n", indent))
	b.WriteString(fmt.Sprintf("%s    current_user,\n", indent))
	b.WriteString(fmt.Sprintf("%s    action=\"%s\",\n", indent, op.Action))
	b.WriteString(fmt.Sprintf("%s    resource=\"%s\",\n", indent, op.Resource))
	writeAuthInputs(b, op.Inputs, indent)
	writeAuthResourceID(b, op, indent)
	b.WriteString(fmt.Sprintf("%s)\n", indent))
}
