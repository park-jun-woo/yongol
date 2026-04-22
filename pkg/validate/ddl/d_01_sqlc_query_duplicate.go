//ff:func feature=validate type=rule control=iteration dimension=2 topic=ddl-structural
//ff:what D-1 — sqlc query name duplicate

package ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d01SqlcQueryDuplicate validates D-1: a `-- name: <Name>` query that appears
// more than once across db/queries/*.sql is an ERROR. Because sqlc uses a global
// namespace, a ModelPrefix is recommended.
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
					"[D-1] sqlc query name %q is duplicated — sqlc uses a global namespace",
					name),
				Advice: fmt.Sprintf("Add a ModelPrefix to make -- name: %s unique (e.g. User%s, Gig%s)", name, name, name),
			})
		}
	}
	return diags
}
