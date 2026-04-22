//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what q09CheckSelectOnExec — detects top-level SELECT / RETURNING in an :exec query and generates an ERROR diagnostic

package query

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

// q09CheckSelectOnExec inspects a single :exec query and returns (diag, true)
// when it emits rows via a top-level SELECT or a RETURNING clause.
func q09CheckSelectOnExec(q sqlc.QuerySpec) (diagnostic.Diagnostic, bool) {
	if q.Cardinality != "exec" {
		return diagnostic.Diagnostic{}, false
	}
	body, err := readQueryBody(q)
	if err != nil || body == nil {
		return diagnostic.Diagnostic{}, false
	}
	text := strings.TrimSpace(body.Text)
	hasTopSelect := topLevelSelectRe.MatchString(text)
	hasReturning := returningWordRe.MatchString(body.Text)
	if !hasTopSelect && !hasReturning {
		return diagnostic.Diagnostic{}, false
	}
	return diagnostic.Diagnostic{
		File:    q.File,
		Line:    q.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[Q-09] :exec query " + q.Name + " returns rows (top-level SELECT / RETURNING present)",
		Advice:  "If rows are needed, change the cardinality to :one / :many; otherwise remove SELECT/RETURNING",
	}, true
}
