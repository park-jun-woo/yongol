//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what checkXqs73Guard — @empty/@exists 시퀀스에서 부분 SELECT PK 컬럼 존재 검사

package ssac_sqlc

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// checkXqs73Guard checks @empty/@exists sequences for partial SELECT field violations.
// These sequences access var.ID (the PK column).
func checkXqs73Guard(fn ssacparser.ServiceFunc, seq ssacparser.Sequence, vars map[string]xqs73VarInfo) []diagnostic.Diagnostic {
	target := seq.Target
	vi, ok := vars[target]
	if !ok {
		return nil
	}
	// @empty and @exists check var.ID — the PK column
	if !selectColsContain(vi.query.SelectCols, "id") {
		return []diagnostic.Diagnostic{{
			File:    fn.FileName,
			Line:    seq.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[XQS-73] @%s %q accesses the PK column (id), but query %q does not SELECT it (partial SELECT: %s)", seq.Type, target, vi.query.Name, strings.Join(vi.query.SelectCols, ", ")),
			Advice:  fmt.Sprintf("Add \"id\" to the SELECT column list of query %q, or use SELECT *.", vi.query.Name),
		}}
	}
	return nil
}
