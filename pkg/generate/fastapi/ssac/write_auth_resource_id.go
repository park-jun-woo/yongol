//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeAuthResourceID — Ownership 유무에 따라 authz_check 의 resource_id/owners 인자 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeAuthResourceID writes the resource_id (and owners, when ownership is
// present) keyword arguments of the authz_check call.
func writeAuthResourceID(b *strings.Builder, op *ir.AuthOp, indent string) {
	if op.Ownership == nil {
		writeResourceIDFromInputs(b, op.Inputs, indent)
		return
	}
	ow := op.Ownership
	b.WriteString(fmt.Sprintf("%s    resource_id=str(%s),\n", indent, ow.ResourcePK))
	b.WriteString(fmt.Sprintf("%s    owners={\"%s\": {\"%s\": owner}},\n",
		indent, op.Resource, ow.OwnerColumn))
}
