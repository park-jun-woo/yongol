//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildEmpty — @empty 시퀀스 빌더 (zero-value guard)

package ssac

import (
	"fmt"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func (g *methodGen) buildEmpty(seq ssacparser.Sequence) ([]string, []string) {
	status := resolveErrStatus("empty", seq.ErrStatus)
	msg := seq.Message
	if msg == "" {
		msg = neutralMessage(status)
	}
	lines := []string{
		fmt.Sprintf("if %s {", zeroValueCheck(seq.Target+".ID")),
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		fmt.Sprintf("\treturn api.%s%dJSONResponse{Error: %q, Code: strPtr(%q)}, nil", g.FuncName, status, msg, neutralCode(status)),
		"}",
	}
	return lines, []string{`"log/slog"`}
}
