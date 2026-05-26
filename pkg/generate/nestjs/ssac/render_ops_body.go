//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what renderOpsBody — Op 배열 → TypeScript 문장 일괄 렌더링 (dispatch)

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderOpsBody writes the TypeScript statements for each Op in the plan.
func renderOpsBody(b *strings.Builder, ops []ir.Op, indent, prismaRef string) {
	for _, op := range ops {
		renderOneOp(b, op, indent, prismaRef)
	}
}
