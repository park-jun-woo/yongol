//ff:func feature=validate type=rule control=iteration dimension=1 topic=query-structural
//ff:what Q-03 — 쿼리명 PascalCase 강제

package query

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// q03NamePascalCase validates Q-03: query names must start with A-Z and
// contain no underscores. sqlc emits query name as Go method name — non
// PascalCase breaks idiom and may collide with other symbols.
func q03NamePascalCase(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, q := range fs.SQLcQueries {
		if q.Name == "" {
			continue
		}
		first := q.Name[0]
		hasUnderscore := strings.ContainsRune(q.Name, '_')
		if first >= 'A' && first <= 'Z' && !hasUnderscore {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    q.File,
			Line:    q.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[Q-03] query name " + q.Name + " is not PascalCase",
			Advice:  "Go 메서드 네이밍 규약을 따르도록 대문자로 시작하고 underscore 없이 작성하세요 (예: GetUser, WorkflowListByOrgID)",
		})
	}
	return diags
}
