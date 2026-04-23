//ff:func feature=gen-gogin type=generator control=sequence
//ff:what renderNextStateFile — <ID>NextState 접근자 함수 단일 Go source

package state

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// renderNextStateFile emits package statemachine source with exactly one
// function — <Symbol>NextState(currentState, event string) string — so
// filefunc F1 (1 file 1 func) passes. See BUG-002 for why symbol
// (PascalCase) is the only form allowed in emitted Go identifiers.
func renderNextStateFile(id, symbol string) string {
	_ = id
	var b strings.Builder
	b.WriteString(ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{
			Feature: "statemachine",
			Type:    "util",
			Control: "sequence",
			Topic:   "state-transition",
		},
		What: symbol + "NextState — 현재 상태 + 이벤트 → 다음 상태 반환 (불가 시 \"\")",
	}))

	b.WriteString("package statemachine\n\n")

	fmt.Fprintf(&b, "// %sNextState returns the target state after a valid transition.\n", symbol)
	b.WriteString("// Returns empty string when the transition is not allowed.\n")
	fmt.Fprintf(&b, "func %sNextState(currentState, event string) string {\n", symbol)
	fmt.Fprintf(&b, "\tevents, ok := %sTransitions[currentState]\n", symbol)
	b.WriteString("\tif !ok {\n\t\treturn \"\"\n\t}\n")
	b.WriteString("\treturn events[event]\n")
	b.WriteString("}\n")

	return b.String()
}
