//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildEval — @eval predicate guard 빌더 (true → early return STATUS)

package ssac

import (
	"fmt"
	"strings"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildEval emits a predicate-guard block for an `@eval` sequence:
//
//	if pkg.Func(pkg.FuncRequest{Field: source}) {
//	    slog.<lvl>("handler: <tag>", "op", "<funcName>", "status", <code>)
//	    return api.<Op><Code>JSONResponse{Error: "<msg>", Code: "<code_name>"}, nil
//	}
//
// `true` triggers the guard (early-return) — the polarity contract documented
// in docs/ssac.md and the Phase001 plan. STATUS is mandatory at the SSaC
// level (S-68); buildEval still defers to resolveErrStatus("eval", ...) so an
// unforeseen 0 in tests/snapshots maps to 500 rather than panicking.
func (g *methodGen) buildEval(seq ssacparser.Sequence) ([]string, []string) {
	parts := strings.SplitN(seq.Model, ".", 2)
	pkgName := parts[0]
	callFunc := ""
	if len(parts) == 2 {
		callFunc = parts[1]
	}
	fields := g.mapFields(seq.Inputs)

	status := resolveErrStatus("eval", seq.ErrStatus)
	msg := seq.Message
	if msg == "" {
		msg = neutralMessage(status)
	}

	lines := []string{
		fmt.Sprintf("if %s.%s(%s.%sRequest{%s}) {", pkgName, callFunc, pkgName, callFunc, fields),
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		"\t" + g.guardReturn(msg, status),
		"}",
	}
	imports := []string{`"log/slog"`}
	imports = append(imports, g.buildEvalImports(seq)...)
	if g.IsSubscribe {
		imports = append(imports, `"fmt"`)
	}
	return lines, imports
}
