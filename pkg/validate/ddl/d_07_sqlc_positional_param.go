//ff:func feature=validate type=rule control=iteration dimension=2 topic=ddl-structural
//ff:what D-7 — positional parameters ($1, $2, …) forbidden in sqlc queries

package ddl

import (
	"fmt"
	"regexp"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

var positionalParamRe = regexp.MustCompile(`\$\d+`)

// d07SqlcPositionalParam validates D-7: $N positional parameters are forbidden.
// Use @name for WHERE/SET/VALUES, sqlc.arg(name) for LIMIT/OFFSET.
func d07SqlcPositionalParam(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, q := range fs.SQLcQueries {
		for _, h := range scanPositionals(q.File, q.Line) {
			diags = append(diags, diagnostic.Diagnostic{
				File:  q.File,
				Line:  h.line,
				Phase: diagnostic.PhaseValidate,
				Level: diagnostic.LevelError,
				Message: fmt.Sprintf("[D-7] query %s — positional parameter %s is forbidden",
					q.Name, h.param),
				Advice: "Use @name for WHERE/SET/VALUES and sqlc.arg(name) for LIMIT/OFFSET",
			})
		}
	}
	return diags
}
