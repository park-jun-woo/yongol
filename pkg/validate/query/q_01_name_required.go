//ff:func feature=validate type=rule control=iteration dimension=1 topic=query-structural
//ff:what Q-01 — the `-- name:` annotation is required

package query

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// q01NameRequired validates Q-01: every non-empty SQL statement in a .sql file
// under db/queries/ must be preceded by a `-- name:` comment. Standalone SQL
// without annotation indicates accidental commits or misplaced DDL.
func q01NameRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	// Collect all query files referenced by QuerySpecs; if none, nothing to scan.
	fileSet := make(map[string]bool)
	for _, q := range fs.SQLcQueries {
		fileSet[q.File] = true
	}
	for file := range fileSet {
		pendingLine, needsName := q01ScanForMissingName(file)
		if !needsName {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    file,
			Line:    pendingLine,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[Q-01] SQL statement in " + filepath.Base(file) + " has no `-- name:` annotation",
			Advice:  "Add a `-- name: <PascalName> :one|:many|:exec` comment to each query",
		})
	}
	return diags
}
