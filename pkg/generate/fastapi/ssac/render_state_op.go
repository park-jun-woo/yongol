//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderStateOp — StateOp → 상태 전이 검증 가드 Python 코드 렌더링

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
	statusExpr := "current_state"
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

	b.WriteString(fmt.Sprintf("%s# @state %s.%s — transition guard\n", indent, op.Diagram, op.Transition))
	b.WriteString(fmt.Sprintf("%sallowed_%s: dict[str, list[str]] = {\n", indent, snakeCase(op.Transition)))
	b.WriteString(fmt.Sprintf("%s    # TODO: populate from Mermaid stateDiagram '%s'\n", indent, op.Diagram))
	b.WriteString(fmt.Sprintf("%s}\n", indent))
	b.WriteString(fmt.Sprintf("%sif \"%s\" not in (allowed_%s.get(%s) or []):\n",
		indent, snakeCase(op.Transition), snakeCase(op.Transition), statusExpr))
	b.WriteString(fmt.Sprintf("%s    raise HTTPException(status_code=%d, detail=\"%s\")\n",
		indent, statusCode, msg))
}
