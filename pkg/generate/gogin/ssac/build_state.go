//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildState — @state 시퀀스 빌더 (상태 전이 검증)

package ssac

import (
	"fmt"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func (g *methodGen) buildState(seq ssacparser.Sequence) ([]string, []string) {
	status := resolveErrStatus("state", seq.ErrStatus)
	msg := seq.Message
	if msg == "" {
		msg = neutralMessage(status)
	}
	inputField := ""
	for _, v := range seq.Inputs {
		inputField = g.mapValue(v)
		break
	}
	// Look up the exported PascalCase Go symbol for this diagram. The
	// DiagramID stays as the lowercase SSaC reference (matching the
	// filename stem), but the generated code must use the Symbol so
	// cross-package callers can reach the function — see BUG-002.
	// Fall back to DiagramID if the lookup is empty (parser always
	// populates it; this keeps codegen resilient to future parser
	// refactors).
	sym := g.DiagramSymbol[seq.DiagramID]
	if sym == "" {
		sym = seq.DiagramID
	}
	lines := []string{
		fmt.Sprintf("if !statemachine.%sCanTransition(%s, %q) {", sym, inputField, seq.Transition),
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		"\t" + g.guardReturn(msg, status),
		"}",
	}
	imports := []string{fmt.Sprintf(`"%s/internal/statemachine"`, g.ModulePath), `"log/slog"`}
	if g.IsSubscribe {
		imports = append(imports, `"fmt"`)
	}
	return lines, imports
}
