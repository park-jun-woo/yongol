//ff:func feature=gen-fastapi type=util control=iteration dimension=1
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

	// Emit ownership lookup when Ownership metadata exists.
	if op.Ownership != nil {
		ow := op.Ownership
		model := pascalCase(ow.Table)
		b.WriteString(fmt.Sprintf("%sowner_row = await session.execute(\n", indent))
		b.WriteString(fmt.Sprintf("%s    select(%s.%s).where(%s.%s == %s)\n",
			indent, model, ow.OwnerColumn, model, ow.ResourcePK, ow.ResourcePK))
		b.WriteString(fmt.Sprintf("%s)\n", indent))
		b.WriteString(fmt.Sprintf("%sowner = owner_row.scalar_one_or_none()\n", indent))
	}

	b.WriteString(fmt.Sprintf("%sawait authz_check(\n", indent))
	b.WriteString(fmt.Sprintf("%s    current_user,\n", indent))
	b.WriteString(fmt.Sprintf("%s    action=\"%s\",\n", indent, op.Action))
	b.WriteString(fmt.Sprintf("%s    resource=\"%s\",\n", indent, op.Resource))
	// Render inputs, skipping ResourceID (handled separately below).
	for _, input := range op.Inputs {
		if input.Key == "ResourceID" {
			continue
		}
		b.WriteString(fmt.Sprintf("%s    %s=%s,\n", indent, resolveArgKey(input), renderArgValue(input)))
	}
	if op.Ownership != nil {
		ow := op.Ownership
		b.WriteString(fmt.Sprintf("%s    resource_id=str(%s),\n", indent, ow.ResourcePK))
		b.WriteString(fmt.Sprintf("%s    owners={\"%s\": {\"%s\": owner}},\n",
			indent, ow.Table, ow.OwnerColumn))
	} else {
		// No ownership but ResourceID may still be present in inputs.
		for _, input := range op.Inputs {
			if input.Key == "ResourceID" {
				b.WriteString(fmt.Sprintf("%s    resource_id=str(%s),\n", indent, renderArgValue(input)))
			}
		}
	}
	b.WriteString(fmt.Sprintf("%s)\n", indent))
}
