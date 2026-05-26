//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderOpsBody — Op 배열 → Python 문장 일괄 렌더링 (dispatch)

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderOpsBody writes the Python statements for each Op in the plan.
func renderOpsBody(b *strings.Builder, ops []ir.Op, indent, sessionRef string) {
	for _, op := range ops {
		renderOneOp(b, op, indent, sessionRef)
	}
}
