//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what renderStateOp — StateOp.AllowedFromStates 기반 상태 전이 검증 가드 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderStateOp writes a state machine transition validation guard. When
// AllowedFromStates is populated (from Phase018 IR Mermaid stateDiagram
// enrichment), actual source states are rendered. Otherwise a TODO comment
// is emitted.
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

	if len(op.AllowedFromStates) > 0 {
		// Render actual allowed states from Mermaid stateDiagram.
		b.WriteString(fmt.Sprintf("%sconst allowed_%s: Record<string, boolean> = {\n", indent, op.Transition))
		for _, state := range op.AllowedFromStates {
			b.WriteString(fmt.Sprintf("%s  '%s': true,\n", indent, state))
		}
		b.WriteString(fmt.Sprintf("%s};\n", indent))
		b.WriteString(fmt.Sprintf("%sif (!allowed_%s[%s]) {\n",
			indent, op.Transition, statusExpr))
	} else {
		// Fallback: no stateDiagram available.
		b.WriteString(fmt.Sprintf("%sconst allowed_%s: Record<string, boolean> = {\n", indent, op.Transition))
		b.WriteString(fmt.Sprintf("%s  // TODO: populate from Mermaid stateDiagram '%s'\n", indent, op.Diagram))
		b.WriteString(fmt.Sprintf("%s};\n", indent))
		b.WriteString(fmt.Sprintf("%sif (!allowed_%s[%s]) {\n",
			indent, op.Transition, statusExpr))
	}

	b.WriteString(fmt.Sprintf("%s  throw new HttpException('%s', HttpStatus.%s);\n",
		indent, msg, httpStatusConst(statusCode)))
	b.WriteString(fmt.Sprintf("%s}\n", indent))
}
