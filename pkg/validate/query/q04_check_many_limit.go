//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what q04CheckManyLimit — checks whether a single :many query has a LIMIT clause and generates a WARNING diagnostic if not

package query

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

// q04CheckManyLimit inspects a single :many query and returns (diag, true)
// when it lacks a LIMIT clause and is not escaped via `-- @no-pagination`.
func q04CheckManyLimit(q sqlc.QuerySpec) (diagnostic.Diagnostic, bool) {
	if q.Cardinality != "many" {
		return diagnostic.Diagnostic{}, false
	}
	body, err := readQueryBody(q)
	if err != nil || body == nil {
		return diagnostic.Diagnostic{}, false
	}
	if body.Escapes["@no-pagination"] {
		return diagnostic.Diagnostic{}, false
	}
	upper := strings.ToUpper(body.Text)
	if strings.Contains(upper, " LIMIT ") || strings.Contains(upper, "\nLIMIT") {
		return diagnostic.Diagnostic{}, false
	}
	return diagnostic.Diagnostic{
		File:    q.File,
		Line:    q.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[Q-04] :many query " + q.Name + " has no LIMIT clause (unbounded result)",
		Advice:  "Add a LIMIT clause, or add `-- +no-pagination` if returning all rows is intentional",
	}, true
}
