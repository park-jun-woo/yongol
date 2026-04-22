//ff:func feature=validate type=rule control=sequence topic=migration-hints
//ff:what Run — MIG-001~006 전체 실행 (generate 파이프라인이 호출)
package migration

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// Run executes every MIG-* rule against the prev/curr schemas and the
// hints. missing is the sidecar-file list from LoadDataMigrationSQL.
// specsDir is required for MIG-006 (snapshot drift check).
func Run(
	specsDir string,
	prev, curr *migration.Schema,
	hints *migration.Hints,
	issues []migration.SafetyIssue,
	missing []string,
) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, Mig001RenameMismatch(prev, curr, hints)...)
	diags = append(diags, Mig002NotNullWithoutBackfill(issues)...)
	diags = append(diags, Mig003DataMigrationMissing(missing)...)
	diags = append(diags, Mig004DestructiveWithoutAllow(issues)...)
	diags = append(diags, Mig005CastMissing(issues)...)
	diags = append(diags, Mig006SnapshotDrift(specsDir)...)
	return diags
}
