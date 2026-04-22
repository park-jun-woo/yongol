//ff:func feature=gen-gogin type=generator control=sequence
//ff:what renderStateFile — ID + transitionMap → Transitions 변수만 담은 Go source (1파일 1decl)

package state

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// renderStateFile assembles the Go source for the transition table of a
// single statemachine. Emits only the <ID>Transitions map variable —
// the CanTransition guard and NextState accessor now live in sibling
// files rendered by renderCanTransitionFile / renderNextStateFile so
// filefunc F1 passes on the statemachine package.
func renderStateFile(id string, transMap map[string]map[string]string) string {
	var b strings.Builder
	b.WriteString(ffannot.EmitAnnotationBlock(ffannot.Block{
		Type: ffannot.TypeAnnot{
			Feature: "statemachine",
			Type:    "model",
		},
		What: id + "Transitions — 상태 전이 테이블 (currentState, event → nextState)",
	}))

	b.WriteString("package statemachine\n\n")

	fmt.Fprintf(&b, "// %sTransitions maps (currentState, event) → nextState.\n", id)
	fmt.Fprintf(&b, "// Generated from states/%s.md — do not edit.\n", id)
	fmt.Fprintf(&b, "var %sTransitions = map[string]map[string]string{\n", id)
	renderTransitionEntries(&b, transMap)
	b.WriteString("}\n")

	return b.String()
}
