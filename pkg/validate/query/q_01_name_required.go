//ff:func feature=validate type=rule control=iteration dimension=1 topic=query-structural
//ff:what Q-01 — `-- name:` 어노테이션 필수

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
			Advice:  "각 쿼리에 `-- name: <PascalName> :one|:many|:exec` 주석을 추가하세요",
		})
	}
	return diags
}
