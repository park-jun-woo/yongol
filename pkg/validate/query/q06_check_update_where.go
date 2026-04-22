//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what q06CheckUpdateWhere — checks whether a single UPDATE query has a WHERE clause and generates an ERROR diagnostic if not

package query

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

// q06CheckUpdateWhere inspects a single query and returns (diag, true) when
// it is an UPDATE without a WHERE clause and not escaped with
// `-- @allow-truncate`.
func q06CheckUpdateWhere(q sqlc.QuerySpec) (diagnostic.Diagnostic, bool) {
	body, err := readQueryBody(q)
	if err != nil || body == nil {
		return diagnostic.Diagnostic{}, false
	}
	text := strings.TrimSpace(body.Text)
	if !updateStartRe.MatchString(text) {
		return diagnostic.Diagnostic{}, false
	}
	if body.Escapes["@allow-truncate"] {
		return diagnostic.Diagnostic{}, false
	}
	upper := strings.ToUpper(text)
	if strings.Contains(upper, " WHERE ") || strings.Contains(upper, "\nWHERE") {
		return diagnostic.Diagnostic{}, false
	}
	return diagnostic.Diagnostic{
		File:    q.File,
		Line:    q.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[Q-06] UPDATE in " + q.Name + " has no WHERE clause (potential full-table update)",
		Advice:  "Add a WHERE clause, or add `-- +allow-truncate` if a full-table update is intentional",
	}, true
}
