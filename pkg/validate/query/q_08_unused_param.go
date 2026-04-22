//ff:func feature=validate type=rule control=iteration dimension=2 topic=query-structural
//ff:what Q-08 — 선언된 파라미터가 본문에 안 쓰이면 ERROR

package query

import (
	"regexp"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// declaredParamRe matches the Params list entries already parsed from the
// SQL body (named params only; positional $N is rejected by D-07).
var paramRefRe = regexp.MustCompile(`(?:@([a-z_][a-z0-9_]*))|(?:sqlc\.arg\(\s*([a-z_][a-z0-9_]*)\s*\))`)

// q08UnusedParam validates Q-08: QuerySpec.Params list includes every
// discovered `@name` / `sqlc.arg(name)` from the SQL body. If the parser
// recorded a parameter yet no reference survives a second scan, the
// parameter is dead (usually after a WHERE clause edit). Keep it or remove.
func q08UnusedParam(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, q := range fs.SQLcQueries {
		if len(q.Params) == 0 {
			continue
		}
		body, err := readQueryBody(q)
		if err != nil || body == nil {
			continue
		}
		found := make(map[string]bool)
		for _, m := range paramRefRe.FindAllStringSubmatch(body.Text, -1) {
			name := m[1]
			if name == "" {
				name = m[2]
			}
			if name != "" {
				found[caseconv.SnakeToPascalSqlc(name)] = true
			}
		}
		for _, declared := range q.Params {
			if found[declared] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    q.File,
				Line:    q.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[Q-08] query " + q.Name + " declares unused parameter " + declared,
				Advice:  "선언된 파라미터 " + declared + " 를 WHERE/SET/VALUES 에 참조하거나 제거하세요",
			})
		}
	}
	return diags
}
