//ff:func feature=validate type=rule control=iteration dimension=1 topic=query-structural
//ff:what Q-03 — query names must be in PascalCase

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
			Advice:  "Follow Go method naming conventions: start with an uppercase letter and omit underscores (e.g. GetUser, WorkflowListByOrgID)",
		})
	}
	return diags
}
