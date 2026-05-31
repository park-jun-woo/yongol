//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeAuthResourceID — Ownership 유무에 따라 authz.check 의 resourceId/owners 인자 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeAuthResourceID writes the resourceId (and owners, when ownership is
// present) entries of the authz.check object.
func writeAuthResourceID(b *strings.Builder, op *ir.AuthOp, indent string) {
	if op.Ownership == nil {
		writeResourceIDFromInputs(b, op.Inputs, indent)
		return
	}
	ow := op.Ownership
	b.WriteString(fmt.Sprintf("%s  resourceId: String(params.%s),\n", indent, ow.ResourcePK))
	b.WriteString(fmt.Sprintf("%s  owners: { %s: { %s: owner?.%s } },\n",
		indent, op.Resource, ow.OwnerColumn, ow.OwnerColumn))
}
