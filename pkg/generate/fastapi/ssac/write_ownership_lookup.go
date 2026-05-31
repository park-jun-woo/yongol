//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeOwnershipLookup — Ownership 메타 기반 owner row SELECT lookup Python 코드 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeOwnershipLookup emits the ownership DB lookup query that fetches the
// owner column for the resource before the authz check.
func writeOwnershipLookup(b *strings.Builder, ow *ir.OwnershipInfo, indent string) {
	model := pascalCase(ir.DDLTableSingularIR(ow.Table))
	b.WriteString(fmt.Sprintf("%sowner_row = await session.execute(\n", indent))
	b.WriteString(fmt.Sprintf("%s    select(%s.%s).where(%s.%s == %s)\n",
		indent, model, ow.OwnerColumn, model, ow.ResourcePK, ow.ResourcePK))
	b.WriteString(fmt.Sprintf("%s)\n", indent))
	b.WriteString(fmt.Sprintf("%sowner = owner_row.scalar_one_or_none()\n", indent))
}
