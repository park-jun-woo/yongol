//ff:func feature=validate type=rule control=iteration dimension=2 topic=ddl-structural
//ff:what D-7 — sqlc 쿼리에 위치 파라미터($1, $2 등) 전면 금지

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
				Message: fmt.Sprintf("[D-7] query %s — 위치 파라미터 %s 사용 금지",
					q.Name, h.param),
				Advice: "WHERE/SET/VALUES 는 @name, LIMIT/OFFSET 은 sqlc.arg(name) 을 사용하세요",
			})
		}
	}
	return diags
}
