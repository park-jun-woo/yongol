//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderExistsOp — ExistsOp → Python 존재 가드 if 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderExistsOp writes a non-null guard that raises HTTPException.
func renderExistsOp(b *strings.Builder, op *ir.ExistsOp, indent string) {
	if op == nil {
		return
	}
	b.WriteString(fmt.Sprintf("%sif %s:\n", indent, op.VarName))
	b.WriteString(fmt.Sprintf("%s    raise HTTPException(status_code=%d, detail=\"%s\")\n",
		indent, op.StatusCode, op.Message))
}
