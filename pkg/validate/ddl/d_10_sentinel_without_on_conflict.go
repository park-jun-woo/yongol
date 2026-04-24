//ff:func feature=validate type=rule control=iteration dimension=2 topic=ddl-structural
//ff:what D-10 — `@sentinel` INSERT must include `ON CONFLICT DO NOTHING`

package ddl

import (
	"fmt"
	"regexp"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d10OnConflictDoNothingRe mirrors the parser's detector.
var d10OnConflictDoNothingRe = regexp.MustCompile(`(?is)ON\s+CONFLICT\b[^;]*\bDO\s+NOTHING\b`)

// d10SentinelWithoutOnConflict validates D-10: every `@sentinel`-tagged
// INSERT must contain `ON CONFLICT ... DO NOTHING` so repeated
// migration application stays safe.
func d10SentinelWithoutOnConflict(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	files := readDBSQLFiles(fs.SpecsDir)
	if len(files) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, f := range files {
		for _, r := range scanInsertsWithAnnotations(f.content) {
			if !r.Annotated {
				continue
			}
			if d10OnConflictDoNothingRe.MatchString(r.SQL) {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:  f.path,
				Line:  r.StartLine,
				Phase: diagnostic.PhaseValidate,
				Level: diagnostic.LevelError,
				Message: fmt.Sprintf(
					"[D-10] `@sentinel` INSERT into %q is missing `ON CONFLICT DO NOTHING`.",
					r.Table),
				Advice: "Add `ON CONFLICT DO NOTHING` to make the sentinel re-run safe across migrations.",
			})
		}
	}
	return diags
}
