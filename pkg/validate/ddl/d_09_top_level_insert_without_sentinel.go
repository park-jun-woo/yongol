//ff:func feature=validate type=rule control=iteration dimension=2 topic=ddl-structural
//ff:what D-9 — top-level INSERT without `-- @sentinel` annotation is forbidden

package ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d09TopLevelInsertWithoutSentinel validates D-9: any top-level INSERT
// inside a DDL file (specs/db/*.sql) must be preceded by a `-- @sentinel`
// annotation, otherwise the statement would be silently dropped by the
// migration emitter.
func d09TopLevelInsertWithoutSentinel(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	files := readDBSQLFiles(fs.SpecsDir)
	if len(files) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, f := range files {
		for _, r := range scanInsertsWithAnnotations(f.content) {
			if r.Annotated {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:  f.path,
				Line:  r.StartLine,
				Phase: diagnostic.PhaseValidate,
				Level: diagnostic.LevelError,
				Message: fmt.Sprintf(
					"[D-9] Top-level INSERT into %q has no `-- @sentinel` annotation.",
					r.Table),
				Advice: "Add `-- @sentinel` directly above the INSERT, or move the INSERT out of db/*.sql.",
			})
		}
	}
	return diags
}
