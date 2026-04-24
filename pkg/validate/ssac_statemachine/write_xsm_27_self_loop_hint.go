//ff:func feature=validate type=rule control=sequence topic=states
//ff:what writeXsm27SelfLoopHint — XSM-27 advice 내 self-loop 전이 힌트 작성

package ssac_statemachine

import (
	"strings"
)

// writeXsm27SelfLoopHint writes the "If <op> is not declared as a
// transition ..." paragraph that suggests an `<initial> --> <initial>:
// <op>` self-loop when the operation does not change state.
func writeXsm27SelfLoopHint(b *strings.Builder, fnName, diagramID, initial string) {
	b.WriteString("  If ")
	b.WriteString(fnName)
	b.WriteString(" is not declared as a transition in states/")
	b.WriteString(diagramID)
	b.WriteString(".md, add it (self-loop if no state change):\n")
	b.WriteString("    ")
	b.WriteString(initial)
	b.WriteString(" --> ")
	b.WriteString(initial)
	b.WriteString(": ")
	b.WriteString(fnName)
	b.WriteString("\n")
}
