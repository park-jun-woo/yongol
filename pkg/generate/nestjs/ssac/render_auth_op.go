//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderAuthOp — AuthOp.Ownership 기반 DB lookup + AuthzService.check 호출 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderAuthOp writes an authz service check call. When Ownership is
// populated (from Phase018 IR Rego enrichment), an ownership DB lookup
// query is emitted before the authz check. prismaRef is "tx" inside a
// transaction or "this.prisma" otherwise.
func renderAuthOp(b *strings.Builder, op *ir.AuthOp, indent, prismaRef string) {
	if op == nil {
		return
	}

	if op.Ownership != nil {
		writeOwnershipLookup(b, op.Ownership, indent, prismaRef)
	}

	b.WriteString(fmt.Sprintf("%sawait this.authz.check({\n", indent))
	b.WriteString(fmt.Sprintf("%s  action: '%s',\n", indent, op.Action))
	b.WriteString(fmt.Sprintf("%s  resource: '%s',\n", indent, op.Resource))
	writeAuthInputs(b, op.Inputs, indent)
	writeAuthResourceID(b, op, indent)
	b.WriteString(fmt.Sprintf("%s});\n", indent))
}
