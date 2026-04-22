//ff:func feature=validate type=rule control=iteration dimension=2 topic=ddl-structural
//ff:what D-1 — sqlc query name 중복

package ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d01SqlcQueryDuplicate validates D-1: db/queries/*.sql 사이에서 동일한
// `-- name: <Name>` 쿼리 이름이 두 번 이상 등장하면 ERROR. sqlc는 전역
// 네임스페이스이므로 ModelPrefix를 권장한다.
func d01SqlcQueryDuplicate(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.SQLcQueries) == 0 {
		return nil
	}

	type occurrence struct {
		file string
		line int
	}
	occurrences := make(map[string][]occurrence)
	for _, q := range fs.SQLcQueries {
		occurrences[q.Name] = append(occurrences[q.Name], occurrence{
			file: q.File,
			line: q.Line,
		})
	}

	var diags []diagnostic.Diagnostic
	for name, occs := range occurrences {
		if len(occs) < 2 {
			continue
		}
		// Emit a diagnostic anchored at every duplicate occurrence so editors
		// can jump to each conflicting site.
		for _, o := range occs {
			diags = append(diags, diagnostic.Diagnostic{
				File:  o.file,
				Line:  o.line,
				Phase: diagnostic.PhaseValidate,
				Level: diagnostic.LevelError,
				Message: fmt.Sprintf(
					"[D-1] sqlc query name %q is duplicated — sqlc 는 전역 네임스페이스",
					name),
				Advice: fmt.Sprintf("-- name: %s 의 이름이 중복되지 않도록 ModelPrefix 를 추가하세요 (예: User%s, Gig%s)", name, name, name),
			})
		}
	}
	return diags
}
