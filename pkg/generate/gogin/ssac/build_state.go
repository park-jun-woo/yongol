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
	lines := []string{
		fmt.Sprintf("if !statemachine.%sCanTransition(%s, %q) {", seq.DiagramID, inputField, seq.Transition),
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		fmt.Sprintf("\treturn api.%s%dJSONResponse{Error: %q, Code: strPtr(%q)}, nil", g.FuncName, status, msg, neutralCode(status)),
		"}",
	}
	return lines, []string{fmt.Sprintf(`"%s/internal/statemachine"`, g.ModulePath), `"log/slog"`}
}
