//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what renderStateOp — StateOp → 상태 전이 검증 가드 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderStateOp writes a state machine transition validation guard. The guard
// checks if the current state allows the requested transition.
func renderStateOp(b *strings.Builder, op *ir.StateOp, indent string) {
	if op == nil {
		return
	}

	// Find the status input to determine the current state variable.
	statusExpr := "currentState"
	for _, input := range op.Inputs {
		if input.Key == "status" || input.Key == "state" {
			statusExpr = renderArgValue(input)
			break
		}
	}

	statusCode := op.StatusCode
	if statusCode == 0 {
		statusCode = 409
	}

	msg := op.Message
	if msg == "" {
		msg = "invalid state transition"
	}

	b.WriteString(fmt.Sprintf("%s// @state %s.%s — transition guard\n", indent, op.Diagram, op.Transition))
	b.WriteString(fmt.Sprintf("%sconst allowed_%s: Record<string, string[]> = {\n", indent, op.Transition))
	b.WriteString(fmt.Sprintf("%s  // TODO: populate from Mermaid stateDiagram '%s'\n", indent, op.Diagram))
	b.WriteString(fmt.Sprintf("%s};\n", indent))
	b.WriteString(fmt.Sprintf("%sif (!(allowed_%s[%s] || []).includes('%s')) {\n",
		indent, op.Transition, statusExpr, op.Transition))
	b.WriteString(fmt.Sprintf("%s  throw new HttpException('%s', HttpStatus.%s);\n",
		indent, msg, httpStatusConst(statusCode)))
	b.WriteString(fmt.Sprintf("%s}\n", indent))
}
