//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildGet — @get 시퀀스 빌더 (sqlc SELECT, @empty/@exists 후속 시 ErrNoRows 관용)

package ssac

import (
	"fmt"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildGet emits the sqlc SELECT call and error handling. When the following
// sequence is @empty or @exists targeting the same variable, pgx.ErrNoRows is
// treated as "zero value, continue" — @empty checks ID == 0 afterwards to
// return 404, @exists checks ID != 0 to return 409 (and skips on absence).
// Without a matching guard sequence, ErrNoRows propagates as 500 (same as before).
//
// Phase005 pgx/v5 refit — sqlc pgx/v5 returns pgx.ErrNoRows (not
// sql.ErrNoRows) on an empty :one select. The error check switches to the
// pgx sentinel; imports follow.
func (g *methodGen) buildGet(seq ssacparser.Sequence, next *ssacparser.Sequence) ([]string, []string) {
	method := resolveSQLCMethod(seq.Model)
	varName := "_"
	if seq.Result != nil {
		varName = seq.Result.Var
	}
	assign := g.assignOp(varName != "_")
	preamble, argStr, argImports := g.sqlcArgs(method, seq.Inputs)

	imports := append([]string(nil), argImports...)
	errHandler := "if err != nil { return nil, err }"
	if varName != "_" && next != nil && (next.Type == "empty" || next.Type == "exists") && next.Target == varName {
		errHandler = "if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }"
		imports = append(imports, `"github.com/jackc/pgx/v5"`, `"errors"`)
	}

	lines := append([]string(nil), preamble...)
	lines = append(lines,
		fmt.Sprintf("%s, err %s %s.%s(%s)", varName, assign, g.queryVar(), method, argStr),
		errHandler,
	)
	return lines, imports
}
