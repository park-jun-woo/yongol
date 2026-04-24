//ff:func feature=validate type=rule control=sequence topic=states
//ff:what writeXsm27OptionA — XSM-27 advice Option A (state-dependent) 라인 작성

package ssac_statemachine

import (
	"strings"
)

// writeXsm27OptionA writes the "Option A (state-dependent)" stanza that
// tells the author how to add a `// @state <diagram> {...} ...` guard
// above the offending SSaC function.
func writeXsm27OptionA(b *strings.Builder, fnName string, target *statefulTarget, diagramID, varName string) {
	b.WriteString("Option A (state-dependent): add above `func ")
	b.WriteString(fnName)
	b.WriteString("() {}`\n")
	b.WriteString("    // @state ")
	b.WriteString(diagramID)
	b.WriteString(" {")
	b.WriteString(pascalCaseFromLower(target.StateColumn))
	b.WriteString(": ")
	b.WriteString(varName)
	b.WriteString(".")
	b.WriteString(pascalCaseFromLower(target.StateColumn))
	b.WriteString("} \"")
	b.WriteString(fnName)
	b.WriteString("\" \"Cannot ")
	b.WriteString(fnName)
	b.WriteString("\" 409\n")
}
