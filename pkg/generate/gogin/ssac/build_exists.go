//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildExists — @exists 시퀀스 빌더 (non-zero guard)

package ssac

import (
	"fmt"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func (g *methodGen) buildExists(seq ssacparser.Sequence) ([]string, []string) {
	status := resolveErrStatus("exists", seq.ErrStatus)
	msg := seq.Message
	if msg == "" {
		msg = neutralMessage(status)
	}
	lines := []string{
		fmt.Sprintf("if %s {", nonZeroCheck(seq.Target+".ID")),
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		fmt.Sprintf("\treturn api.%s%dJSONResponse{Error: %q, Code: strPtr(%q)}, nil", g.FuncName, status, msg, neutralCode(status)),
		"}",
	}
	return lines, []string{`"log/slog"`}
}
