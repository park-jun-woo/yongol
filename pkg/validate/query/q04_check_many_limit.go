//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what q04CheckManyLimit — 단일 :many 쿼리의 LIMIT 유무 판정 + WARNING 진단 생성

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
		Advice:  "LIMIT 을 추가하거나 의도적 전체 반환이면 `-- +no-pagination` 주석을 붙이세요",
	}, true
}
