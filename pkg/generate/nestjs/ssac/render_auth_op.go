//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what renderAuthOp — AuthOp.Ownership 기반 DB lookup + AuthzService.check 호출 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderAuthOp writes an authz service check call. When Ownership is
// populated (from Phase018 IR Rego enrichment), an ownership DB lookup
// query is emitted before the authz check.
func renderAuthOp(b *strings.Builder, op *ir.AuthOp, indent string) {
	if op == nil {
		return
	}

	// Emit ownership lookup when Ownership metadata exists.
	if op.Ownership != nil {
		ow := op.Ownership
		model := lcFirst(ow.Table)
		b.WriteString(fmt.Sprintf("%sconst owner = await tx.%s.findUnique({\n", indent, model))
		b.WriteString(fmt.Sprintf("%s  where: { %s: params.%s },\n", indent, ow.ResourcePK, ow.ResourcePK))
		b.WriteString(fmt.Sprintf("%s  select: { %s: true },\n", indent, ow.OwnerColumn))
		b.WriteString(fmt.Sprintf("%s});\n", indent))
	}

	b.WriteString(fmt.Sprintf("%sawait this.authz.check({\n", indent))
	b.WriteString(fmt.Sprintf("%s  action: '%s',\n", indent, op.Action))
	b.WriteString(fmt.Sprintf("%s  resource: '%s',\n", indent, op.Resource))
	for _, input := range op.Inputs {
		b.WriteString(fmt.Sprintf("%s  %s: %s,\n", indent, input.Key, renderArgValue(input)))
	}
	if op.Ownership != nil {
		ow := op.Ownership
		b.WriteString(fmt.Sprintf("%s  resourceId: String(params.%s),\n", indent, ow.ResourcePK))
		b.WriteString(fmt.Sprintf("%s  owners: { %s: { %s: owner?.%s } },\n",
			indent, ow.Table, ow.OwnerColumn, ow.OwnerColumn))
	}
	b.WriteString(fmt.Sprintf("%s});\n", indent))
}
