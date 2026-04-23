//ff:func feature=gen-gogin type=generator control=sequence
//ff:what renderCanTransitionFile — <ID>CanTransition 가드 함수 단일 Go source

package state

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// renderCanTransitionFile emits package statemachine source with exactly
// one function — <Symbol>CanTransition(currentState, event string) bool — so
// filefunc F1 (1 file 1 func) passes. See BUG-002: id stays lowercase
// (filename stem); only symbol (PascalCase) appears in generated Go
// code so callers in other packages can reference the function.
func renderCanTransitionFile(id, symbol string) string {
	_ = id
	var b strings.Builder
	b.WriteString(ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{
			Feature: "statemachine",
			Type:    "util",
			Control: "sequence",
			Topic:   "state-transition",
		},
		What: symbol + "CanTransition — 현재 상태에서 이벤트 전이 가능 여부 반환",
	}))

	b.WriteString("package statemachine\n\n")

	fmt.Fprintf(&b, "// %sCanTransition returns true when event is a valid transition\n", symbol)
	fmt.Fprintf(&b, "// from currentState.\n")
	fmt.Fprintf(&b, "func %sCanTransition(currentState, event string) bool {\n", symbol)
	fmt.Fprintf(&b, "\tevents, ok := %sTransitions[currentState]\n", symbol)
	b.WriteString("\tif !ok {\n\t\treturn false\n\t}\n")
	b.WriteString("\t_, ok = events[event]\n")
	b.WriteString("\treturn ok\n")
	b.WriteString("}\n")

	return b.String()
}
