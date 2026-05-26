//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderStateOp — StateOp → 상태 전이 검증 TODO 주석 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderStateOp writes a state machine validation placeholder comment.
func renderStateOp(b *strings.Builder, op *ir.StateOp, indent string) {
	if op == nil {
		return
	}
	b.WriteString(fmt.Sprintf("%s// @state %s.%s\n", indent, op.Diagram, op.Transition))
	b.WriteString(fmt.Sprintf("%s// TODO: integrate state machine validation\n", indent))
}
