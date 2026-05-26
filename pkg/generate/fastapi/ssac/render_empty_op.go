//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderEmptyOp — EmptyOp → Python not 가드 if 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderEmptyOp writes a null guard that raises HTTPException.
func renderEmptyOp(b *strings.Builder, op *ir.EmptyOp, indent string) {
	if op == nil {
		return
	}
	b.WriteString(fmt.Sprintf("%sif not %s:\n", indent, op.VarName))
	b.WriteString(fmt.Sprintf("%s    raise HTTPException(status_code=%d, detail=\"%s\")\n",
		indent, op.StatusCode, op.Message))
}
