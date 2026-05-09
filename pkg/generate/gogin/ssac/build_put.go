//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildPut — @put 시퀀스 빌더 (sqlc UPDATE :exec)

package ssac

import (
	"fmt"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func (g *methodGen) buildPut(seq ssacparser.Sequence) ([]string, []string) {
	method := resolveSQLCMethod(seq.Model)
	preamble, argStr, argImports := g.sqlcArgs(method, seq.Inputs)
	// When the preamble declared a fresh err, the subsequent sqlc call
	// reuses that err via `=`. When no preamble is emitted (no JSONB
	// input), fall back to the original firstErr toggle so the first
	// err-only line still uses `:=` and later ones use `=`.
	var assign string
	if len(preamble) > 0 {
		g.FirstErr = false
		assign = "="
	} else {
		assign = g.assignOp(false)
	}
	lines := append([]string(nil), preamble...)
	lines = append(lines,
		fmt.Sprintf("err %s %s.%s(%s)", assign, g.queryVar(), method, argStr),
		"if err != nil { " + g.returnErr() + " }",
	)
	return lines, argImports
}
