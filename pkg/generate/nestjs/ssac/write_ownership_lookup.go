//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeOwnershipLookup — Ownership 메타 기반 owner findUnique lookup TypeScript 코드 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeOwnershipLookup emits the ownership DB lookup that fetches the owner
// column for the resource before the authz check. prismaRef is "tx" inside a
// transaction or "this.prisma" otherwise.
func writeOwnershipLookup(b *strings.Builder, ow *ir.OwnershipInfo, indent, prismaRef string) {
	// Use singular Prisma model name (DDL table is plural).
	model := lcFirst(ir.DDLTableSingularIR(ow.Table))
	b.WriteString(fmt.Sprintf("%sconst owner = await %s.%s.findUnique({\n", indent, prismaRef, model))
	b.WriteString(fmt.Sprintf("%s  where: { %s: params.%s },\n", indent, ow.ResourcePK, ow.ResourcePK))
	b.WriteString(fmt.Sprintf("%s  select: { %s: true },\n", indent, ow.OwnerColumn))
	b.WriteString(fmt.Sprintf("%s});\n", indent))
}
