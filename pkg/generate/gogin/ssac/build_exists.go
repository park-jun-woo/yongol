//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildExists — @exists 시퀀스 빌더 (non-zero guard, pgtypex NilCheck 분기)

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
	col := g.lookupPKColumn(seq.Target)
	guard := nonZeroCheckWithCol(seq.Target+".ID", col)
	imports := []string{`"log/slog"`}
	if col != nil {
		imports = append(imports, pgtypexImportIfNeeded(col)...)
	}
	lines := []string{
		fmt.Sprintf("if %s {", guard),
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		"\t" + g.guardReturn(msg, status),
		"}",
	}
	if g.IsSubscribe {
		imports = append(imports, `"fmt"`)
	}
	return lines, imports
}
